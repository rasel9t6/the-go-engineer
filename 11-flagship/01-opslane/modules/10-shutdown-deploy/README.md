# OPSL.10 Graceful Shutdown and Deployment

## Mission

Finish Opslane as a production-shaped service that can stop, recover, and deploy safely.

## What This Module Builds

- explicit shutdown handling
- drain behavior for requests and workers
- deployment packaging and health semantics
- final integrated proof for the required flagship path

## You Are Here If

- `OPSL.9` is complete
- the service has background work and observability
- the final concern is safe operation under deploy and shutdown pressure

## Proof Surface

This module is not implemented in the current tree yet.

When it is complete:

- `go build` must succeed for the Opslane server
- the full Opslane test suite must stay green
- shutdown drills must prove drain behavior under in-flight work

Expected new files:

- `cmd/server/shutdown.go`
- any supporting worker drain and readiness logic required by the final design

Existing deployment surfaces that should still participate:

- `Dockerfile`
- `docker-compose.yml`
- repository root `.github/workflows/ci.yml`

## Required Files and Boundaries

Shutdown behavior must coordinate HTTP, database, and background work. This is the final integrated proof, not a one-file patch.

## Engineering Questions

- What should the health endpoint return during drain?
- What guarantees do we make about in-flight requests and queued work?
- Which operational promises are enforceable, and which are only best-effort?

## Next Module

This completes the required Opslane flagship path.
