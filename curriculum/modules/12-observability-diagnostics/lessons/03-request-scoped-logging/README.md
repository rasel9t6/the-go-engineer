# Request-scoped logging

## Mission

Understand and apply Request-scoped logging in the context of professional Go software engineering.

## Prerequisites

- core-12-02

## Mental Model

Request-scoped logging is like putting a label on every file in a project folder. Each request gets a folder (context) with a unique label (correlation_id). Every log line from that request goes into the folder with the label automatically stamped on it. After the request completes, you can search for the label and find every log line in chronological order — a complete trace of what happened during that request. Without request scoping, log lines from 1000 concurrent requests are interleaved in a single stream with no way to group them by request.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

slog.FromContext(ctx) is implemented in log/sdk/context.go. It retrieves a *slog.Logger from context using a private context key. If no logger is found, it returns slog.Default(). slog.WithContext(ctx, logger) stores a logger in context using the same private key. The context-based log functions (InfoContext, ErrorContext) call FromContext internally, so they automatically use the stored logger. When a handler is created via logger.With(attrs...), the original handler is wrapped in a new handler that prepends the stored attrs to every Record before calling Handle. This means attrs added via With are evaluated per-handler, not per-Record — they share a single attr slice to avoid allocation.

## Run Instructions

```bash
go run ./curriculum/modules/12-observability-diagnostics/lessons/03-request-scoped-logging
go test ./curriculum/modules/12-observability-diagnostics/lessons/03-request-scoped-logging
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using context.Background() instead of the request context — if you create a logger with context.Background(), it loses all request-scoped fields (correlation ID, user ID, trace ID). The logger must be derived from the request's context.Context so that WithContext attaches fields to the active request.
- Attaching fields in middleware but logging before the middleware runs — if the logger is created at handler registration time (package init) rather than per-request, all requests use the same logger with no request-specific fields. The logger must be created or derived inside the middleware, not before it.
- Storing a *slog.Logger in a global variable instead of the request context — a global logger cannot have per-request fields. Store *slog.Logger in context.Context using a custom context key, and retrieve it in handlers via slog.FromContext(ctx) (Go 1.21+).

## In Production

Request-scoped logging is universal in production Go services. Every HTTP framework (chi, gin, echo) provides middleware for request-scoped loggers. In observability platforms (Datadog, Grafana, Honeycomb), the correlation_id is the primary dimension for grouping log lines, traces, and metrics into a single 'request' view. Without request scoping, each incident investigation starts with 'find the correlation_id for that failed request' — and you cannot find it if it was never logged.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-12-04`.
