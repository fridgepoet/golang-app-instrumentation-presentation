package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

var tracer = otel.Tracer("")
var logger = otelslog.NewLogger("")

func githubHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "entering the githubHandler")
	defer span.End()

	waitAround(ctx)

	req, err := http.NewRequestWithContext(r.Context(), "GET", "https://api.github.com/", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Add("Accept", `application/json`)

	client := http.Client{
		Timeout:   time.Duration(1) * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rand.Int()%2 != 0 {
		span.SetStatus(codes.Error, "github handler failed")
		span.RecordError(fmt.Errorf("looks like you got unlucky"))
		logger.ErrorContext(ctx, "this logs the made up error")
		http.Error(w, "this is a made up error", http.StatusInternalServerError)
		return
	}

	fmt.Printf("client.Do response status : %s \n", resp.Status)

	fmt.Fprint(w, "External call ok\n")
}

func waitAround(context context.Context) {
	_, span := tracer.Start(context, "starting to wait around")
	defer span.End()

	// Add sleep to make spans more distinct
	time.Sleep(200 * time.Millisecond)
}
