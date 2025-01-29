package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func setupOTelSDK(ctx context.Context) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}

	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	conn, err := initConn()
	if err != nil {
		handleErr(err)
		return
	}

	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	tracerProvider, err := newTraceProvider(ctx, conn)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)
	otel.SetTracerProvider(tracerProvider)

	// Set up logger provider.
	loggerProvider, err := newLoggerProvider(ctx, conn)
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs, loggerProvider.Shutdown)
	global.SetLoggerProvider(loggerProvider)

	return
}

// Initialize a gRPC connection to be used by both the tracer and meter
// providers.
func initConn() (*grpc.ClientConn, error) {
	// It connects the OpenTelemetry Collector through local gRPC connection.
	// You may replace `localhost:4317` with your endpoint.
	conn, err := grpc.NewClient("otel-collector:4317",
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	return conn, err
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTraceProvider(ctx context.Context, conn *grpc.ClientConn) (*trace.TracerProvider, error) {
	//Set up a trace exporter
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter, trace.WithBatchTimeout(1*time.Minute)),
		trace.WithResource(initResource()),
	)
	return traceProvider, nil
}

func newLoggerProvider(ctx context.Context, conn *grpc.ClientConn) (*log.LoggerProvider, error) {
	logExporter, err := otlploggrpc.New(ctx, otlploggrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, err
	}

	log.NewLoggerProvider()
	loggerProvider := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter)),
	)
	return loggerProvider, nil
}

var thisResource *resource.Resource
var initResourcesOnce sync.Once

func initResource() *resource.Resource {
	initResourcesOnce.Do(func() {
		extraResources, err := resource.New(
			context.Background(),
			resource.WithOS(),
			resource.WithProcess(),
			resource.WithContainer(),
			resource.WithHost(),
			resource.WithAttributes(
				semconv.ServiceName("golang-app"),
			),
		)
		if err != nil {
			panic(err)
		}

		thisResource, err = resource.Merge(
			resource.Default(),
			extraResources,
		)
		if err != nil {
			panic(err)
		}
	})
	return thisResource
}
