# Retries and idempotent consumers

## Mission

Understand and apply Retries and idempotent consumers in the context of professional Go software engineering.

## Prerequisites

- core-14-10

## Mental Model

Retries handle transient failures (network blips, timeouts, 503s). Idempotency ensures that retries do not cause duplicate side effects. They are a pair: retries increase reliability, idempotency prevents retries from causing harm. Exponential backoff ensures that retries do not amplify load on an already-degraded system. Circuit breakers stop retrying when the downstream is known to be down — retrying against a dead service is worse than not retrying at all.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Idempotency is typically implemented with a conditional write: INSERT INTO idempotency_keys (key, response, created_at) VALUES ($1, NULL, NOW()) ON CONFLICT (key) DO NOTHING RETURNING created_at. If the INSERT succeeds (no conflict), this is the first request — proceed with processing. If the INSERT returns a row (conflict), this is a retry — return the cached response. The key must be unique and have a TTL (24 hours is typical for payment APIs). For distributed systems, the idempotency store must be strongly consistent (single-node Redis with fsync, or PostgreSQL with synchronous replication) — an eventually-consistent store can accept duplicate writes during a network partition.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/11-retries-and-idempotent-consumers
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/11-retries-and-idempotent-consumers
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Retrying without exponential backoff — retrying every 100ms against a failing database amplifies load by 10x per second, turning a transient failure into a self-inflicted outage.
- Assuming all errors are retryable — a 400 Bad Request retried 5 times is 5 wasted calls; only 5xx, timeouts, and network errors should be retried.
- Not implementing idempotency keys — a payment webhook is delivered twice; without an idempotency key, the customer is charged twice and the error is invisible until the customer notices.
- Retrying inside a transaction that has already rolled back — the retried operation runs outside the transaction context, creating inconsistent state between the rollback and the retry.
- Using a fixed retry count without a circuit breaker — a downstream service is down for 10 minutes, every caller retries 3 times with 1-second backoff, generating 1800 failed calls instead of 600.

## In Production

Every production Go service that calls external APIs implements retries with exponential backoff and jitter. Stripe's API requires idempotency keys for all POST requests — every payment integration must implement this. Kafka consumers use idempotent processing (the consumer tracks the offset of each message; processing the same offset twice is a no-op). AWS SDK for Go implements automatic retries with exponential backoff for all service calls.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-12`.
