## Run app
1. `docker-compose up`
Changes are live reloaded with https://github.com/air-verse/air.

## Make requests
* `curl http://localhost:8080/github-api`

## Try out autoinstrumentation
Check out the [`beyla` branch](https://github.com/fridgepoet/golang-app-instrumentation-presentation/tree/beyla).

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
