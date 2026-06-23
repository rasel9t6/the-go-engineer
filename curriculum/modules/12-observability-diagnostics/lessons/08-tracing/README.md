# Tracing

## Mission

Understand and apply Tracing in the context of professional Go software engineering.

## Prerequisites

- core-12-07

## Mental Model

Tracing is like a detailed itinerary of a package delivery. Each stop (span) records: what happened (HTTP request), who handled it (Service B), how long it took (8ms), and what was done (SELECT orders). The trace is the complete itinerary from the client's first request to the final response, showing every stop and the time spent at each. A critical feature: the itinerary shows NOT just that service B took 8ms, but that 3ms of that was the database query — the slow component is visible. Tracing answers 'where is the time going?' — the question that metrics (aggregates) and logs (individual events) cannot answer for a specific request.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

OpenTelemetry tracing in Go works through several components. TracerProvider creates Tracer instances. Tracer.Start(ctx, name) creates a Span. Start first extracts the parent span from the context (via context.Value with a private key). If no parent exists, the span is a root. The new span receives a new span_id (8 random bytes) and inherits the parent's trace_id. Span.End() records the end timestamp, computes duration, and sends the span to the registered SpanProcessor (typically a BatchSpanProcessor that queues spans and exports them in batches). The exporter serializes spans as OTLP protobuf and sends them to the backend via gRPC or HTTP. The trace context is propagated using W3C TraceContext: the traceparent header contains trace_id (16 bytes hex), parent_span_id (8 bytes hex), and trace_flags (2 hex chars). OpenTelemetry's propagation package provides TextMapPropagator that injects/extracts this header from HTTP request/response.

## Run Instructions

```bash
go run ./curriculum/modules/12-observability-diagnostics/lessons/08-tracing
go test ./curriculum/modules/12-observability-diagnostics/lessons/08-tracing
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Not propagating the trace context through the call chain — if middleware starts a span but passes context.Background() to the handler instead of the span's context, the handler creates a new root trace with no parent — the trace is fragmented into disconnected spans. Always pass r.Context() (which includes the trace context) to downstream calls.
- Creating too many spans — adding a span for every function call (including trivial getters) creates thousands of spans per request, overwhelming the trace backend and making the trace waterfall unreadable. Add spans at service boundaries (HTTP handler → service → DB → external API), not at every function call.
- Sampling too aggressively in production — sampling 100% of requests generates 1000 traces/second on a 1000 req/s service, costing significant CPU and memory. Use probabilistic sampling (1% of requests) combined with rate-limited sampling to keep trace volume manageable while ensuring slow and error traces are captured.

## In Production

Distributed tracing is essential for microservice architectures. Kubernetes services use OpenTelemetry for tracing. AWS X-Ray traces requests across Lambda, API Gateway, and DynamoDB. Datadog APM traces requests across Go services with automatic instrumentation. In incident response, the trace waterfall is the first place operators look: it immediately shows which service hop is slow or erroring.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-12-09`.
