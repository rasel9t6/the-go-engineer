# Why observability exists

## Mission

Understand and apply Why observability exists in the context of professional Go software engineering.

## Prerequisites

- core-11-28

## Mental Model

Observability is the property of a system that lets you understand its internal state from the outside, without shipping new code. If you need to add a log line and redeploy to debug an incident, your system is not observable. Logs tell you what happened, metrics tell you how many times it happened, and traces tell you exactly where it happened in a distributed call graph. All three must be designed in from the start — retrofitting observability during an incident is too late.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

log/slog's Logger holds a handler chain. The default text handler writes to stderr. The JSON handler writes a JSON object per line with time, level, message, and any Attr key-value pairs. The context handler (slog.NewLogLogger) extracts attributes from context and includes them in every log line. Prometheus client uses a registry of collectors (Counter, Gauge, Histogram, Summary). Each collector is lock-free using atomic operations for counters and gauges, and a concurrent map + mutex for histograms. OpenTelemetry's trace provider creates spans with a trace ID, span ID, parent span ID, start time, end time, and attributes. Spans are exported to a collector (e.g., Jaeger, Zipkin) via gRPC or HTTP.

## Run Instructions

```bash
go run ./curriculum/modules/12-observability-diagnostics/lessons/01-why-observability-exists
go test ./curriculum/modules/12-observability-diagnostics/lessons/01-why-observability-exists
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Adding logging only after a failure is detected — observability is treated as an afterthought rather than designed in from the start, leaving the team blind during the incident.
- Logging everything at INFO level — log volume drowns out actionable signals, and the team cannot distinguish routine operations from failure indicators.
- Using fmt.Println for production diagnostics — unstructured text cannot be filtered, aggregated, or alerted on; every incident requires grep through raw log files.
- Monitoring only CPU and memory — the application can be at 10% CPU with a deadlocked goroutine pool, and no alert fires because the infrastructure metrics look healthy.
- Assuming structured logging is enough — logs without correlation IDs, trace IDs, and request-scoped context cannot be linked across services during a multi-service incident.

## In Production

Every production Go service at scale depends on observability. Kubernetes uses structured logging throughout, exposing metrics via the metrics endpoint. Docker Hub uses correlation IDs across its microservices. Uber's observability pipeline processes billions of spans per day. Without observability, a 10-query database slowdown is invisible until users report errors — with observability, the p99 latency chart shows the degradation in real time.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-12-02`.
