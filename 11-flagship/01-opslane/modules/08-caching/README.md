# OPSL.8 Caching Layer

## Mission

Introduce cache-aware behavior without pretending cache invalidation is free.

## What This Module Builds

- cache storage boundary
- cache-aside reads
- TTL behavior
- invalidation rules around changing order and payment state

## You Are Here If

- `OPSL.7` is complete
- there are read paths worth caching
- asynchronous work exists that can make stale data more dangerous

## Proof Surface

This module is not implemented in the current tree yet.

When it is complete:

- `go test` must pass for the future `11-flagship/01-opslane/internal/cache` package
- invalidation tests must prove cache updates follow writes intentionally

Expected new files:

- `internal/cache/cache.go`
- `internal/cache/store.go`
- `internal/middleware/cache.go`

## Required Files and Boundaries

Caching should be additive, not a hidden source of truth.
PostgreSQL remains the system of record.

## Engineering Questions

- Which reads are safe to cache?
- What should invalidate cached order or payment data?
- How will you prevent stampedes when hot keys expire?

## Next Module

Continue to [OPSL.9 Observability](../09-observability/README.md).
