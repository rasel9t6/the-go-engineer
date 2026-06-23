# Logs

## Mission

Understand and apply Logs in the context of professional Go software engineering.

## Prerequisites

- core-15-02

## Mental Model

Logs are the black box flight recorder for your service. Just as pilots review flight data after an incident, operators query logs to reconstruct request flows, measure latency, and correlate events across services. Each log line is a timestamped sensor reading: structured fields (altitude, speed, heading) let you filter for anomalies; unstructured chatter ("flaps deployed") is noise. The aggregation pipeline is the investigation board — if your sensors don't speak a common format (JSON), the board is useless.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}) returns a handler that encodes each log entry as a JSON object. The handler's Handle() method is called for every enabled log event; it calls encoder.Encode() which writes to the io.Writer. Under the hood, sync.Pool reuses encoder buffers to reduce allocations. The Level type is an int; HandlerOptions.Level is a Leveler interface that can be a dynamic atomic.Int64 for live log level changes. When slog.Info() is called, the logger acquires the handler's mutex, calls Handle(), and returns — this is why synchronous I/O blocks the caller. For async behavior, wrap the handler in a custom writer that sends entries to a channel consumed by a background goroutine.

## Run Instructions

```bash
go run ./curriculum/modules/15-docker-cicd-deployment/lessons/03-logs
go test ./curriculum/modules/15-docker-cicd-deployment/lessons/03-logs
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Logging to files instead of stdout: Beginners write logs to /var/log/app.log, violating the 12-factor app principle. In Docker, container stdout/stderr is captured by the runtime; writing to files fills the container layer and breaks log aggregation.
- Using log.Println for everything: New Go devs use log.Println at every execution point without log levels. Production systems need DEBUG, INFO, WARN, ERROR, FATAL levels so operators can filter noise and route errors to alerting pipelines.
- Calling log.Fatal in library code: Learners call log.Fatal(err) inside packages, which calls os.Exit(1) and prevents deferred cleanup, closes the process without HTTP 500 responses, and stops the entire container instead of returning an error up the call chain.
- Ignoring structured logging: Writing fmt.Printf("User %s logged in", user) instead of slog.Info("user login", "user_id", id, "ip", remoteAddr). Unstructured logs cannot be parsed by log aggregators (Loki, Elastic, Datadog) for search, dashboards, or alert rules.
- Logging sensitive data: Learners log request bodies, auth tokens, or SQL queries verbatim. Production pipelines ship logs off-cluster; secrets in logs become compliance violations (PCI, HIPAA, SOC2) and credential leaks.

## In Production

Production Go services at Uber, Datadog, and Grafana Labs use structured logging with slog or zerolog. Kubernetes-native Go apps write JSON logs to stdout; these are collected by fluentd or vector, shipped to Loki or Elasticsearch, and queried via LogQL or Kibana. Common patterns: request-scoped loggers with context.Context (slog.With), trace_id propagation for distributed tracing correlation, sampling in high-volume endpoints (health checks, metrics), and panic handlers that log the stack trace before os.Exit(1).

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-15-04`.
