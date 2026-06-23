# Secure logging

## Mission

Understand and apply secure logging in the context of professional Go software engineering.

## Prerequisites

- core-10-23

## Mental Model

Secure logging ensures that sensitive data (passwords, tokens, PII) is never written to log output while still preserving enough context for debugging and monitoring.

## Visual Model

```text
log event -> redaction filter -> structured output -> log aggregator
```

## Machine View

Structured loggers serialize log entries as JSON objects with fields like {"time": "...", "level": "INFO", "msg": "...", "request_id": "...", "key": "value"}. This format is machine-parseable by log aggregation systems. Redaction is implemented as a middleware that scans for known sensitive field names and replaces their values with a redacted sentinel.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/24-secure-logging
go test ./curriculum/modules/10-auth-security/lessons/24-secure-logging
```

## Try It

Add a redaction middleware that scrubs password fields from structured log output.

## In Production

Production systems must never log raw credentials, session tokens, or PII. Use structured logging with redaction middleware and audit log trails for compliance.

## Thinking Questions

- What fields should be redacted by default in a web application?
- How does structured logging improve security incident investigation?

## Next Step

Refresh tokens
