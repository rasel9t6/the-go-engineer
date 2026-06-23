# Metrics

## Mission

Understand and apply Metrics in the context of professional Go software engineering.

## Prerequisites

- core-12-05

## Mental Model

Metrics are the vital signs of a service. Like heart rate (request rate), blood pressure (latency), temperature (error rate), and oxygen level (saturation). Each vital sign is a time series: a sequence of (timestamp, value) pairs. A single number tells you nothing; the trend over time tells you everything. A spike in heart rate (request rate) during a code deploy is expected; a spike in temperature (error rate) is not. Metrics answer 'is the service healthy?' and 'what changed?' — they do not answer 'why?' (that is logs) or 'where exactly?' (that is traces).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Prometheus client_golang stores metrics in a registry (prometheus.Registry). Each metric is a Collector that returns its current values when Collect(ch chan<- Metric) is called. Counter stores a uint64 with an atomic add operation. Histogram stores a slice of uint64 buckets and a total count (via a sync.Mutex for concurrent safety). When the /metrics endpoint is scraped, promhttp.Handler calls Gather() on the registry, which calls Collect on every registered Collector, serializes each Metric's value in Prometheus text format. The text format is: # HELP metric_name description, # TYPE metric_name counter, metric_name{label="value"} 42. Prometheus server scrapes this text, parses the label-value pairs, and stores the time series.

## Run Instructions

```bash
go run ./curriculum/modules/12-observability-diagnostics/lessons/06-metrics
go test ./curriculum/modules/12-observability-diagnostics/lessons/06-metrics
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using gauge for a monotonically increasing value — gauge can go up and down (CPU, memory usage). Counter only goes up (request count, bytes sent). Using gauge for request count means a reset or rollback resets the counter to 0, making rate calculations incorrect. Always use Counter for cumulative totals.
- Forgetting to label metrics by meaningful dimensions — a single metric 'http_requests_total' with no labels tells you the total request count but not which endpoint, method, or status. Without labels, a 5xx spike on /api/login is invisible in the aggregate. Always add labels (endpoint, method, status_code) to differentiate failure modes.
- Exporting metrics from the request path — calling metric.Inc() on every request blocks the handler if the metric backend is slow or the mutex is contended. Use an atomic counter or a buffered channel for hot-path metrics; export the aggregated values on a separate goroutine.

## In Production

Metrics are the foundation of production observability. Every Go service exports Prometheus metrics. Kubernetes uses metrics for horizontal pod autoscaling (HPA) based on CPU/memory and custom metrics. Incident response starts with metrics: 'is this a latency spike, error rate spike, or traffic surge?' — each requires a different mitigation. SRE teams define SLOs (service level objectives) based on metrics (latency < 200ms for 99.9% of requests over 30 days).

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-12-07`.
