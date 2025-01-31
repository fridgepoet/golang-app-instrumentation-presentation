# Getting Started with OpenTelemetry
This repo serves as an exploratory point for instrumenting a Go app and visualizing telemetry data in Grafana Cloud.

## First steps
1. If you want to start by seeing traces in your logs, navigate to OpenTelemetry SDK setup and substitute in an stdout exporter, for example: `traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())`. See an earlier [commit](https://github.com/fridgepoet/golang-app-instrumentation-presentation/blob/a0f83dc64b5d5b612f4419a952e29a4a4ce7ab1c/otel.go#L52-L53).
2. Once you're ready to explore visualizing, you could start exporting to an OTLP endpoint. If you want to export to Grafana Cloud's OTLP endpoint, [see instructions below](#optional-get-grafana-cloud-access-tokens-otlp-endpoint-etc).
Start bringing in the grpc exporters. See otel.go in the `main` branch, for example check out how the trace exporter is set up [here](https://github.com/fridgepoet/golang-app-instrumentation-presentation/blob/e168c91d7da3fac1c2d995f2d44bb27564d46943/otel.go#L90).

## Run app, generate activity, visualize
1. `docker-compose up`
Changes to the golang code are live reloaded with https://github.com/air-verse/air.
2. `curl http://localhost:8080/github-api` a few times to create activity. You have a 1 in 2 chance of hitting a fake error
3. If you're visualizing with Grafana Cloud, navigate to your instance. My favorite way to visualize data is: 
   1. left sidebar > Explore
   2. I choose `grafanacloud-shirley-traces` as my datasource (Tempo is the database for trace storage)
   3. Click on "Search" for the Query type, this lets me see any traces that have come in.
   4. Clicking on a trace ID lets you visualize the spans and attributes.
   5. Explore logs and metrics by switching the datasource. Or check out the curated exploration apps: left sidebar > Explore > below Explore, you'll see Metrics, Logs, Traces...

## (Optional) Get Grafana Cloud access tokens, OTLP endpoint, etc
1. [Create a free Grafana Cloud account](https://grafana.com/pricing/). Go to grafana.com > My account.
2. Click on your stack's name.
3. Look for the OpenTelemetry box. Click Configure.
4. You'll find:
   * URL for the OTLP endpoint
   * the Grafana Cloud instance ID
   * You can generate a Password/API token
5. Use these in the config placeholders in this repo.

Based on these docs: https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/#manual-opentelemetry-setup-for-advanced-users

## Try out auto-instrumentation
Check out the [`beyla` branch](https://github.com/fridgepoet/golang-app-instrumentation-presentation/tree/beyla).

## Links
* [Getting started with OpenTelemetry (Go)](https://opentelemetry.io/docs/languages/go/getting-started/)
* [Beyla (auto-instrumentation)](https://grafana.com/docs/beyla/latest/)
* [Sign up for a free Grafana Cloud account](https://grafana.com/pricing/)
