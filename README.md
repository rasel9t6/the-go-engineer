# The Go Engineer

[![CI](https://github.com/rasel9t6/the-go-engineer/actions/workflows/ci.yml/badge.svg)](https://github.com/rasel9t6/the-go-engineer/actions)
[![License: TGE License v1.0 (Non-Commercial)](https://img.shields.io/badge/License-TGE_v1.0-red.svg?style=for-the-badge)](#license)
[![GitHub Sponsors](https://img.shields.io/badge/sponsor-30363D?style=for-the-badge&logo=GitHub-Sponsors&logoColor=%23EA4AAA)](https://github.com/sponsors/rasel9t6)
[![Patreon](https://img.shields.io/badge/Patreon-F96854?style=for-the-badge&logo=patreon&logoColor=white)](https://patreon.com/rasel9t6)

Learn Go by building, testing, and operating real software.

The Go Engineer is a repo-first Go software engineering learning system built around a progressive **5-phase, 12-section architecture**.

It takes a learner from machine fundamentals and first execution all the way to a production-shaped Go backend capstone.

For the curriculum source of truth, read [ARCHITECTURE.md](./ARCHITECTURE.md).

## Current Status

The v2.1 stable release is shipped.

- `main`: post-v2.1 implementation line
- `release/v1`: stable v1 maintenance line
- `release/v2`: stable v2.1.x maintenance line

Architecture v2.1 is locked. Normal work should deepen, polish, test, document, or maintain the existing 12-section structure. Do not introduce new public root sections unless the maintainer explicitly reopens architecture.

## Quick Start

```bash
git clone https://github.com/rasel9t6/the-go-engineer.git
cd the-go-engineer
go mod download
go version
go run ./00-how-computers-work/1-what-is-a-program
```

## Curriculum Overview

```text
Phase 0: Machine Foundation     s00  How Computers Work          (0% -> 5%)
Phase 1: Language Foundation    s01  Getting Started             (5% -> 12%)
                                s02  Language Basics             (12% -> 28%)
                                s03  Functions & Errors          (28% -> 38%)
                                s04  Types & Design              (38% -> 52%)
Phase 2: Engineering Core       s05  Packages, I/O & CLI         (52% -> 62%)
                                s06  Backend, APIs & Databases   (62% -> 75%)
                                s07  Concurrency                 (75% -> 83%)
                                s08  Quality & Testing           (83% -> 87%)
Phase 3: Systems Engineering    s09  Architecture & Security     (87% -> 92%)
                                s10  Production Operations       (92% -> 96%)
Phase 4: Flagship Project       s11  Opslane SaaS Backend        (96% -> 100%)
```

## Source Sections

| Section | Folder | Focus |
| --- | --- | --- |
| s00 | [00-how-computers-work](./00-how-computers-work) | execution, memory, terminal, processes |
| s01 | [01-getting-started](./01-getting-started) | install, first run, toolchain basics |
| s02 | [02-language-basics](./02-language-basics) | values, control flow, data structures |
| s03 | [03-functions-errors](./03-functions-errors) | reusable logic and explicit failure handling |
| s04 | [04-types-design](./04-types-design) | structs, interfaces, composition, strings and text |
| s05 | [05-packages-io](./05-packages-io) | modules, packages, CLI, encoding, filesystem |
| s06 | [06-backend-db](./06-backend-db) | HTTP servers, API design, gRPC, databases |
| s07 | [07-concurrency](./07-concurrency) | goroutines, context, sync, pipelines |
| s08 | [08-quality-test](./08-quality-test) | tests, profiling, benchmarks |
| s09 | [09-architecture](./09-architecture) | package design, architecture patterns, security |
| s10 | [10-production](./10-production) | logging, shutdown, config, observability, deployment |
| s11 | [11-flagship](./11-flagship) | integrated Opslane capstone system |

## What You Will Build

You will work through:

- beginner-friendly exercises and starter projects
- parsers, filesystem tools, and CLI utilities
- HTTP services and database-backed applications
- concurrency pipelines, worker pools, and timeout-aware clients
- profiling, testing, and benchmark-driven improvements
- structured logging, graceful shutdown, configuration, observability, and deployment workflows
- **Opslane**: a production-shaped SaaS backend capstone

## Core Docs

| Document | Purpose |
| --- | --- |
| [ARCHITECTURE.md](./ARCHITECTURE.md) | curriculum source of truth |
| [CURRICULUM-BLUEPRINT.md](./CURRICULUM-BLUEPRINT.md) | lesson and milestone teaching contract |
| [LEARNING-PATH.md](./LEARNING-PATH.md) | entry points and path guidance |
| [ROADMAP.md](./ROADMAP.md) | post-v2.1 focus, maintenance status, and implementation priorities |
| [docs/PROGRESSION.md](./docs/PROGRESSION.md) | visual progression through phases and milestones |
| [CODE-STANDARDS.md](./CODE-STANDARDS.md) | code and explanation standards |
| [TESTING-STANDARDS.md](./TESTING-STANDARDS.md) | testing patterns and verification expectations |
| [COMMON-MISTAKES.md](./COMMON-MISTAKES.md) | common Go bugs and fixes |
| [CONTRIBUTING.md](./CONTRIBUTING.md) | contribution workflow |
| [RELEASE.md](./RELEASE.md) | release and stabilization process |
| [MAINTAINER-CHECKLIST.md](./MAINTAINER-CHECKLIST.md) | maintainer review and release checklist |

## Repository Workflow

All repository changes should follow the same lightweight operating contract:

1. Create or confirm an approved issue.
2. Branch from the correct long-lived line.
3. Open a draft PR early.
4. Keep commits logical and use the bracketed commit style.
5. Run local verification before review.
6. Wait for maintainer review; do not self-merge.

Commit format:

```text
[TYPE] short imperative description
```

Allowed types:

```text
[FEAT] [FIX] [DOCS] [TEST] [CHORE] [REFACTOR] [SECURITY] [RELEASE]
```

## Validate the Repo

Run the CI-equivalent checks before final review:

```bash
go build ./...
go vet ./...
unformatted=$(gofmt -l .); test -z "$unformatted" || (echo "$unformatted" && exit 1)
go mod tidy
git diff --exit-code -- go.mod go.sum
go test ./...
go test -race ./...
go test -coverprofile=coverage.out ./...
go run ./scripts/validate_curriculum.go
```

For benchmark-related changes:

```bash
go test -bench=. -benchmem -count=1 ./08-quality-test/01-quality-and-performance/testing/benchmarks/
```

## Toolchain Policy

This repository intentionally targets the Go version declared in `go.mod` and CI. If the local toolchain is older than the repo target, do not downgrade the module. Install the required Go version or report the mismatch separately.

## Windows Note

Some database and capstone paths use `go-sqlite3`, which requires CGO and a C compiler. On Windows, WSL2 is the smoothest setup for those paths.

## License

This project is licensed under **The Go Engineer License (TGE License) v1.0**.

- Free for personal, educational, and non-commercial use
- Commercial use requires permission
