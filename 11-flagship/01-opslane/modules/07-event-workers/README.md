# OPSL.7 Event Bus and Worker Pools

## Mission

Add bounded asynchronous work without turning background processing into invisible chaos.

## What This Module Builds

- event bus primitives
- worker pool boundaries
- queueing and backpressure behavior
- safe drain rules for background work

## You Are Here If

- `OPSL.6` is complete
- payment and order workflows are ready to emit asynchronous follow-up work
- you are prepared to reason about bounded concurrency

## Proof Surface

This module is not implemented in the current tree yet.

When it is complete:

- `go test` must pass for the future `internal/events` package
- `go test` must pass for the future `internal/workers` package
- saturation and drain tests must prove bounded behavior

Expected new files:

- `internal/events/bus.go`
- `internal/events/types.go`
- `internal/workers/pool.go`
- `internal/workers/order_processor.go`
- `internal/workers/payment_processor.go`
- `internal/workers/notification_worker.go`

## Required Files and Boundaries

Worker coordination should stay explicit and bounded.
Avoid untracked goroutine spawning inside handlers or repositories.

## Engineering Questions

- What happens when the queue is full?
- When should the system reject work instead of buffering more?
- How do you inspect or drain a stuck worker safely?

## Next Module

Continue to [OPSL.8 Caching Layer](../08-caching/README.md).
