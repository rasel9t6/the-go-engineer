# Correlation IDs

## Mission

Understand and apply Correlation IDs in the context of professional Go software engineering.

## Prerequisites

- core-12-03

## Mental Model

A correlation ID is a single UUID that travels with a request across every service boundary — like a unique ticket number attached to a customer's journey through a multi-department process. Every service that touches the request logs the correlation ID alongside its own events. When debugging a production issue, the operator searches the log aggregator for the correlation ID and sees every log line from every service in chronological order — a complete timeline of the request.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Correlation ID propagation in Go depends on context.Context. The context carries the correlation ID through the call chain without explicit function parameters. The key is an unexported type to prevent key collisions between packages. For HTTP propagation, the standard header is X-Correlation-Id (or X-Request-Id). For gRPC, it is carried in metadata (key-value pairs attached to the gRPC context). For message queues, it is carried in the message headers (e.g., Kafka headers, RabbitMQ message properties). Go's net/http Transport does not automatically propagate context — the caller must explicitly copy the correlation ID from context to the outgoing request header. Tools like OpenTelemetry automate this with instrumentation libraries.

## Run Instructions

```bash
go run ./curriculum/modules/12-observability-diagnostics/lessons/04-correlation-ids
go test ./curriculum/modules/12-observability-diagnostics/lessons/04-correlation-ids
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Generating a new correlation ID at each service boundary instead of propagating the incoming one — the trace is broken into disconnected segments that cannot be correlated during debugging.
- Storing the correlation ID in a global variable or a singleton — concurrent requests overwrite each other's correlation IDs, mixing traces and making every request appear to have the same ID.
- Logging the correlation ID in only some log lines and not others — searching for a trace produces partial results, and the missing log lines cannot be correlated to any request.
- Not forwarding the correlation ID to downstream services in HTTP headers or gRPC metadata — the trace stops at the first service boundary, and downstream operations appear as orphan events.
- Using in-memory-only correlation ID storage — when the request fails, the correlation ID is lost with the process and cannot be recovered from logs or external storage.

## In Production

Correlation IDs are the foundation of observability in distributed systems. Every major production Go service uses them: Kubernetes audit logs, Docker Hub API traces, Uber's observability pipeline. In incident response, the first question is always: 'what is the correlation ID?' Without it, finding the root cause of a multi-service failure requires manually stitching together timestamps from 5 different log dashboards — a 10-minute investigation becomes a 2-hour slog.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-12-05`.
