# OPSL.9 Observability

## Mission

Make Opslane explain itself under load, failure, and change.

## What This Module Builds

- structured logging surfaces
- correlation IDs
- metrics
- trace-friendly request and background-work context

## You Are Here If

- `OPSL.8` is complete
- background work and cache behavior now exist
- debugging the system requires better evidence than raw logs

## Proof Surface

This module is not implemented in the current tree yet.

When it is complete:

- `go test` must pass for the future `internal/logging` package
- `go test` must pass for the future `internal/metrics` package
- correlation tests must prove one request can be followed across layers

Expected new files:

- `internal/logging/logger.go`
- `internal/logging/middleware.go`
- `internal/metrics/metrics.go`
- `internal/tracing/tracing.go`

## Required Files and Boundaries

Instrumentation should help answer why the system is degrading, not just generate more text.

## Engineering Questions

- Which metrics would warn about degradation before users complain?
- What is the minimum log context needed to follow one failed order?
- How should tracing connect HTTP requests to background work?

## Next Module

Continue to [OPSL.10 Graceful Shutdown and Deployment](../10-shutdown-deploy/README.md).
