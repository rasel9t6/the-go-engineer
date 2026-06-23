# Secure logging

## Mission

Understand and apply Secure logging in the context of professional Go software engineering.

## Prerequisites

- core-10-23

## Mental Model

Secure logging is like a medical record: it must contain enough information to be useful for diagnosis but must never contain the patient's full SSN or password. Structured logs are like a spreadsheet: each field is a column, making the data queryable. Unstructured logs are like a novel: you need to read everything to find one fact.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Structured loggers serialize log entries as JSON objects with fields like {"time": "...", "level": "INFO", "msg": "...", "request_id": "...", "key": "value"}. This format is machine-parseable by log aggregation systems. Redaction is implemented as a middleware that scans for known sensitive field names and replaces their values with a redacted sentinel.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/24-secure-logging
go test ./curriculum/modules/10-auth-security/lessons/24-secure-logging
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Logging sensitive data: passwords, tokens, API keys, PII in plaintext.
- Using fmt.Println or log.Println for structured logging — unstructured logs cannot be queried effectively.
- Logging in the wrong severity level: logging everything as INFO makes ERRORs invisible.
- Not sanitizing user input before logging — log injection attacks can corrupt log aggregation systems.

## In Production

Every production Go service uses structured logging. slog is the standard library choice. Zap (uber-go/zap) and Logrus (sirupsen/logrus) are popular third-party alternatives. Datadog, Grafana Loki, and Elasticsearch are common log aggregation backends.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-25`.
