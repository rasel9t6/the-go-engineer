# Structured logging with slog

## Mission

Understand and apply Structured logging with slog in the context of professional Go software engineering.

## Prerequisites

- core-12-01

## Mental Model

Structured logging is the difference between a sticky note ('error occurred') and a spreadsheet row ('time: 12:00:00, level: ERROR, msg: db query failed, query_time_ms: 5200, path: /api/users'). The sticky note is human-readable in isolation but useless for search, aggregation, and alerting. The spreadsheet row is machine-parseable: a log aggregator can sum query_time_ms > 1000, alert on ERROR level, or filter by path. slog's handler API is a factory assembly line: each log call creates a Record (the item on the line), passes it through the handler (the worker that serializes it), and writes it to output (the shipping dock).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

slog.Handler is an interface with three methods: Enabled(ctx, level) returns whether this level is active; Handle(ctx, record) serializes and writes the record; WithAttrs(attrs) returns a new handler with pre-attached attributes. The default JSON handler (slog.NewJSONHandler) implements Handle by writing a JSON object per line. It iterates record.Attrs (a lazily-evaluated attr list), serializes each key-value pair, and appends to a buffer. The variadic log functions (Info, Error, etc.) are convenience wrappers: they call the internal logger.Log which creates a slog.Record, attaches the args as attrs (using reflection to flip pairs), and passes the record to the handler. The Attr variants (LogAttrs, InfoAttrs) skip the reflection step. slog.Record is designed to be allocation-friendly: it stores attrs in a ring buffer (up to 10 inline) and only allocates when exceeded.

## Run Instructions

```bash
go run ./curriculum/modules/12-observability-diagnostics/lessons/02-structured-logging-with-slog
go test ./curriculum/modules/12-observability-diagnostics/lessons/02-structured-logging-with-slog
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Logging sensitive data (passwords, API keys, PII) in plain text — slog has no built-in redaction. If you log r.FormValue("password"), the password appears in every log aggregator. Always redact or hash sensitive fields before passing them to slog.LogAttrs.
- Using slog.Default() in library code — libraries should accept a *slog.Logger parameter (or use context-based logger retrieval) rather than calling slog.Default(), which changes the global logger's behavior for all consumers.
- Ignoring slog's handler API when a custom format is needed — instead of writing a custom handler that implements slog.Handler, developers often parse and re-format log output with regex, which breaks structured fields and level filtering. Implementing slog.Handler is ~50 lines and gives full control.

## In Production

slog is the standard structured logging library in Go 1.21+. Kubernetes uses JSON logging natively. Datadog, Grafana Loki, and AWS CloudWatch all ingest structured JSON logs. In production Go services, slog replaces log.Println with log/slog for every log line — the structured format enables real-time dashboards (error rate by endpoint), historical analysis (latency p99 trend), and automated alerts (error count > threshold).

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-12-03`.
