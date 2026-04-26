# OPSL.6 Payment Pipeline

## Mission

Model payment work as a reliability problem, not just another insert statement.

## What This Module Builds

- payment service workflow
- gateway integration seam
- retry and timeout behavior
- reconciliation-safe duplicate protection

## You Are Here If

- `OPSL.5` is complete
- order state transitions exist
- payment handling can move into a dedicated workflow layer

## Proof Surface

This module is not implemented in the current tree yet.

When it is complete:

- `go test` must pass for the future `internal/payment` package
- payment workflow tests must prove retries and duplicate protection
- reconciliation tests must cover missing callback and timeout scenarios

Expected new files:

- `internal/payment/gateway.go`
- `internal/payment/worker.go`
- `internal/services/payment.go`

## Required Files and Boundaries

Keep gateway behavior behind an explicit boundary.
Do not scatter payment retry logic across handlers and repositories.

## Engineering Questions

- What happens on gateway timeout?
- How do you stop duplicate payment attempts from becoming duplicate charges?
- What should happen when success is observed locally but the follow-up callback never arrives?

## Next Module

Continue to [OPSL.7 Event Bus and Worker Pools](../07-event-workers/README.md).
