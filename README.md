This branch contains the golang app and autoinstrumentation alone, based on https://grafana.com/docs/beyla/latest/setup/docker/.

## Run app, generate activity, visualize
1. `docker-compose up`
2. `curl http://localhost:8080/github-api` a few times to create activity. You have a 1 in 2 chance of hitting a fake error
3. If you're visualizing with Grafana Cloud, navigate to your instance. My favorite way to visualize data is:
   1. left sidebar > Explore
   2. I choose `grafanacloud-shirley-traces` as my datasource (Tempo is the database for trace storage)
   3. Click on "Search" for the Query type, this lets me see any traces that have come in.
   4. Clicking on a trace ID lets you visualize the spans and attributes.
   5. Explore logs and metrics by switching the datasource. Or check out the curated exploration apps: left sidebar > Explore > below Explore, you'll see Metrics, Logs, Traces...

## (Optional) Get Grafana Cloud access tokens, OTLP endpoint, etc

2. [Create a free Grafana Cloud account](https://grafana.com/pricing/). Go to grafana.com > My account.
2. Click on your stack's name.
3. Look for the OpenTelemetry box. Click Configure.
4. You'll find:
   * URL for the OTLP endpoint
   * the Grafana Cloud instance ID
   * You can generate a Password/API token
5. Use these in the config placeholders in this repo.

Based on these docs: https://grafana.com/docs/grafana-cloud/send-data/otlp/send-data-otlp/#manual-opentelemetry-setup-for-advanced-users

## Links
* [Getting started with OpenTelemetry (Go)](https://opentelemetry.io/docs/languages/go/getting-started/)
* [Beyla (auto-instrumentation)](https://grafana.com/docs/beyla/latest/)
* [Sign up for a free Grafana Cloud account](https://grafana.com/pricing/)
