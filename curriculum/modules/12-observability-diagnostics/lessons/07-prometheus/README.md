# Prometheus

## Mission

Understand and apply Prometheus in the context of professional Go software engineering.

## Prerequisites

- core-12-06

## Mental Model

Prometheus is like a meter reader walking a route every 15 minutes (scrape), reading each meter (HTTP endpoint) and recording the numbers. The readings are stored in a ledger (TSDB). An analyst (Grafana) can ask questions about the data (PromQL): 'what was the average request rate over the last hour?' 'what was the p99 latency yesterday at 3pm?' If a meter is unreadable (service down), the reader notes it (target DOWN), and an alarm (Alertmanager) notifies the owner. The key design choice: Prometheus pulls, it does not push — it decides when to read, not the service.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Prometheus TSDB stores each time series as a sequence of (timestamp, value) pairs. Data is written in blocks of 2 hours (configurable). Each block contains a directory with: index (mapping label names/values to series IDs), chunks (compressed sample data), metadata, and tombstones (for deletion markers). The index uses a postings list for fast label matching: finding all series with method="GET" is an O(1) lookup in the postings list. PromQL queries are executed by PromQL engine which parses the query into an AST, optimizes it (pushes down label matchers), and iterates over matching series. The rate() function computes per-second increase over a time range: (last_value - first_value) / (last_timestamp - first_timestamp) × (samples_per_second / scrape_interval). histogram_quantile() estimates the N-th percentile from histogram bucket counts by linear interpolation within the bucket containing the N-th percentile value.

## Run Instructions

```bash
go run ./curriculum/modules/12-observability-diagnostics/lessons/07-prometheus
go test ./curriculum/modules/12-observability-diagnostics/lessons/07-prometheus
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Including process metrics without considering cardinality — promauto auto-registers process metrics (CPU, memory, goroutines, file descriptors). In Kubernetes with hundreds of pod restarts, each pod's process metrics create unique time series — if a pod restarts 10 times, Prometheus stores 10 distinct series for the same metric. Use relabeling rules in Prometheus to drop unnecessary process metrics.
- Not setting a write timeout on the /metrics handler — generating the metrics page for a service with 10K metrics takes seconds. If the handler has no timeout, a slow client connection ties up the goroutine. Fix: use http.Server.ReadTimeout and WriteTimeout on the metrics mux, or set a context deadline on the request.
- Using promhttp.HandlerFor with a registry that panics on duplicate registration — if two metrics have the same name, prometheus.MustRegister panics. In production, use prometheus.NewRegistry() with promhttp.HandlerFor(registry, promhttp.HandlerOpts{ErrorHandling: promhttp.ContinueOnError}) to handle registration conflicts gracefully.

## In Production

Prometheus is the CNCF-graduated standard for metrics monitoring. Kubernetes uses Prometheus metrics for everything: kube-state-metrics exposes cluster state, node_exporter exposes host metrics, and application metrics are scraped via pod annotations. Grafana's default datasource is Prometheus. Almost every open-source Go project exports Prometheus metrics. Prometheus is installed in 80%+ of Kubernetes production clusters.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-12-08`.
