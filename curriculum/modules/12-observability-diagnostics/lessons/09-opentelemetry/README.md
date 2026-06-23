# OpenTelemetry

## Mission

Understand and apply OpenTelemetry in the context of professional Go software engineering.

## Prerequisites

- core-12-08

## Mental Model

OpenTelemetry is a universal adapter for observability. Instead of each service using a different SDK (Jaeger for tracing, Datadog for metrics, Splunk for logs), OTel provides a single API that all services use. The backend (Jaeger, Tempo, Datadog) is plugged in at startup via exporters. This is like USB-C for observability: one connector (OTel API) that works with any device (backend). The OTel Collector is a USB hub: it receives data from many services and routes it to the right backends. The key insight: OTel decouples instrumentation from backend — you can switch from Jaeger to Datadog by changing one line in main.go, not every instrumentation site.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

OpenTelemetry Go SDK consists of several layers. At the bottom is the Exporter (e.g., otlptracegrpc) that implements spanspan.Exporter with ExportSpans(ctx, spans) method. Above that is SpanProcessor: BatchSpanProcessor receives spans on a channel, batches them (up to 512 spans or 1s), and calls exporter.ExportSpans. Above the processor is TracerProvider.Start which creates a Span. The Span struct stores: name, start/end timestamps, parent span context, attributes (a slice of key-value pairs), events, status, and resource. When Span.End() is called, it calls processor.OnEnd(span) which enqueues the span. The BatchSpanProcessor's goroutine dequeues, serializes as OTLP protobuf, and sends via gRPC. The OTLP protobuf schema defines TraceService/Export for traces, MetricService/Export for metrics, and LogsService/Export for logs. The OpenTelemetry Collector receives OTLP, processes it (filter, sample, add attributes), and forwards to one or more backends.

## Run Instructions

```bash
go run ./curriculum/modules/12-observability-diagnostics/lessons/09-opentelemetry
go test ./curriculum/modules/12-observability-diagnostics/lessons/09-opentelemetry
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using the OpenTelemetry SDK directly instead of the API — importing go.opentelemetry.io/otel/sdk instead of go.opentelemetry.io/otel locks your code to a specific SDK implementation. Always depend on the API package and inject the SDK at startup. This lets you swap backends (Jaeger → Datadog → Tempo) without changing instrumentation code.
- Creating a new tracer provider per request — TracerProvider initialization is expensive (reads environment variables, configures exporters, starts goroutines for batch processing). Creating one per request leaks memory and goroutines. Initialize once at startup and store it globally or pass it via dependency injection.
- Not setting the global tracer provider — without otel.SetTracerProvider(provider), the default no-op provider is used, and all spans are silently dropped. The SDK must be explicitly installed as the global provider before any instrumentation runs.

## In Production

OpenTelemetry is the industry standard for observability instrumentation. CNCF adopted OTel as the successor to OpenTracing and OpenCensus. Kubernetes uses OTel for tracing. AWS Distro for OpenTelemetry provides a managed OTel collector. Grafana Tempo, Datadog, Honeycomb, and New Relic all accept OTLP. In production Go, OTel is replacing vendor-specific SDKs: a service instrumented with OTel can be observed by any OTel-compatible backend without code changes.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-12-10`.
