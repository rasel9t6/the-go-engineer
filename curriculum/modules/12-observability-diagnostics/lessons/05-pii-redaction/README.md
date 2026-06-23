# PII redaction

## Mission

Understand and apply PII redaction in the context of professional Go software engineering.

## Prerequisites

- core-12-04

## Mental Model

PII redaction is like a security guard at the logging factory gate. Every log record (the item on the assembly line) passes through the guard before being packaged for shipment. The guard checks each field against a list of 'sensitive' labels (email, password, SSN). If a field's label matches, the guard replaces the value with a masked version (alice@***). The guard never sees the raw value — it only sees the labeled package. This means PII is removed at the source, before the log line leaves the process.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A custom slog.Handler that redacts PII implements the slog.Handler interface: Enabled(ctx, level) delegates to inner handler; Handle(ctx, record) is the core — it creates a copy of the record, iterates record.Attrs (a lazy attr iterator), checks each attr's key against the PII set, and creates a new attr with the redacted value for matches (using slog.String(attr.Key, redacted)). The modified attrs are stored in a new record via slog.NewRecord(time, level, msg, pc) and slog.Record.AddAttrs. The original attrs that were not PII are re-added. The modified record is passed to inner.Handle. This pattern ensures the redacted log has the same structure and all non-PII fields intact.

## Run Instructions

```bash
go run ./curriculum/modules/12-observability-diagnostics/lessons/05-pii-redaction
go test ./curriculum/modules/12-observability-diagnostics/lessons/05-pii-redaction
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Redacting at the presentation layer instead of the logging layer — applying redaction only in the HTTP response (masking email in the JSON body) but not in the log line means the full email appears in logs where operators and third-party aggregators can see it. Redact at every output boundary: response, log, error message, and metrics tag.
- Using regex-based redaction on structured fields — applying a regex like s/(\w+)@/***@/ on the serialized JSON string can match keys instead of values (the JSON key "email" also contains '@'). Apply redaction on the structured value before serialization, not on the serialized blob.
- Assuming redaction once means redacted forever — a code change that adds a new log line with a new user field bypasses existing redaction logic unless redaction is enforced at the handler level (custom slog.Handler) rather than at individual log call sites.

## In Production

PII redaction is legally required for GDPR (EU), CCPA (California), and SOC 2 compliance. Every production Go service that logs user data needs PII redaction. Stripe's logging pipeline redacts credit card numbers before they reach the centralized log store. GitHub's audit log redacts email addresses. In healthcare (HIPAA), PII redaction in logs is mandatory.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-12-06`.
