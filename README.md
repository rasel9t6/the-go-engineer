# The Go Engineer

[![CI](https://github.com/rasel9t6/the-go-engineer/actions/workflows/ci.yml/badge.svg)](https://github.com/rasel9t6/the-go-engineer/actions)
[![License: TGE License v1.0 (Non-Commercial)](https://img.shields.io/badge/License-TGE_v1.0-red.svg?style=for-the-badge&logo=data:image/svg%2Bxml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0id2hpdGUiPjxwYXRoIGQ9Ik0xNCAySDZhMiAyIDAgMCAwLTIgMnYxNmEyIDIgMCAwIDAgMiAyaDEyYTIgMiAwIDAgMCAyLTJWOGwtNi02em0tMSAxLjVMMTguNSA5SDEzVjMuNXpNNiAyMFY0aDV2Nmg2djEwSDZ6bTItOGg4djJIOHYtMnptMCA0aDh2Mkg4di0yeiIgZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiLz48L3N2Zz4%3D)](#license)
[![GitHub Sponsors](https://img.shields.io/badge/sponsor-30363D?style=for-the-badge&logo=GitHub-Sponsors&logoColor=%23EA4AAA)](https://github.com/sponsors/rasel9t6)
[![Patreon](https://img.shields.io/badge/Patreon-F96854?style=for-the-badge&logo=patreon&logoColor=white)](https://patreon.com/rasel9t6)

Learn Go by building, testing, and operating real software.

The Go Engineer is a repo-first Go software engineering learning system built around the beta
11-stage architecture.

That public stage model is now matched by the main learner-facing source surfaces in this repo, so
learners can follow the curriculum by stage without translating between competing folder systems.

If you want the public explanation of that transition model, read
[docs/beta-public-architecture.md](./docs/beta-public-architecture.md).

## Start Here

Pick the release channel that matches what you want:

- `release/v1`: the stable v1 line for learners who want the current supported experience
- `release/v2`: the current public alpha checkpoint for the v2 line
- `main`: the active beta implementation branch and the fastest-moving line
- `v2.0.0-alpha.1`: the first public v2 alpha tag

If you want the safest path today, use `release/v1`.
If you want the current public v2 checkpoint, use `release/v2` or the `v2.0.0-alpha.1` tag.
If you want to follow the beta rollout as it lands, use `main`.

## Quick Start

```bash
git clone https://github.com/rasel9t6/the-go-engineer.git
cd the-go-engineer
go mod download
go version
go run ./01-getting-started/2-hello-world
```

## Public Curriculum (11 Stages)

The curriculum is organized into 11 stages. Each stage has a dedicated entry page under [docs/stages](./docs/stages/README.md).

| Stage | Focus | Sections | Source |
| --- | --- | --- | --- |
| [01 Getting Started](./docs/stages/01-getting-started.md) | install, hello world, dev setup | stage 01 | [01-getting-started](./01-getting-started/) |
| [02 Language Basics](./docs/stages/02-language-basics.md) | variables, types, control flow, data structures | stage 02 | [02-language-basics](./02-language-basics/) |
| [03 Functions & Errors](./docs/stages/03-functions-errors.md) | functions, params, error handling | stage 03 | [03-functions-errors](./03-functions-errors/) |
| [04 Types & Design](./docs/stages/04-types-design.md) | structs, interfaces, composition, strings | stage 04 | [04-types-design](./04-types-design/), [05-composition](./05-composition/), [06-strings-and-text](./06-strings-and-text/) |
| [05 Packages & IO](./docs/stages/05-packages-io.md) | modules, packages, CLI, files | stage 05 | [05-packages-io/01-modules-and-packages](./05-packages-io/01-modules-and-packages/), [05-packages-io/02-io-and-cli](./05-packages-io/02-io-and-cli/) |
| [06 Backend & DB](./docs/stages/06-backend-db.md) | HTTP, databases, handlers | stage 06 | [06-backend-db/01-web-and-database](./06-backend-db/01-web-and-database/) |
| [07 Concurrency](./docs/stages/07-concurrency.md) | goroutines, channels, patterns | stage 07 | [07-concurrency/01-concurrency](./07-concurrency/01-concurrency/), [07-concurrency/02-concurrency-patterns](./07-concurrency/02-concurrency-patterns/) |
| [08 Quality & Test](./docs/stages/08-quality-test.md) | testing, benchmarking, profiling | stage 08 | [08-quality-test/01-quality-and-performance](./08-quality-test/01-quality-and-performance/) |
| [09 Architecture](./docs/stages/09-architecture.md) | package design, services | stage 09 | [09-architecture/01-package-design](./09-architecture/01-package-design/), [09-architecture/02-grpc](./09-architecture/02-grpc/) |
| [10 Production](./docs/stages/10-production.md) | logging, shutdown, deployment | stage 10 | [10-production/01-structured-logging](./10-production/01-structured-logging/), [10-production/02-graceful-shutdown](./10-production/02-graceful-shutdown/), [10-production/03-docker-and-deployment](./10-production/03-docker-and-deployment/) |
| [11 Flagship](./docs/stages/11-flagship.md) | full GoScale project | stage 11 | [11-flagship/01-enterprise-capstone](./11-flagship/01-enterprise-capstone/), [11-flagship/02-code-generation](./11-flagship/02-code-generation/) |

## Best Entry Point By Learner Type

### Complete beginner

Start at [Stage 01: Getting Started](./docs/stages/01-getting-started.md).
That stage page points to the current source surface for beginner setup and first-run work.

### Experienced programmer new to Go

Start at [Stage 02: Language Basics](./docs/stages/02-language-basics.md).
If you want a faster ramp, skim [Stage 01](./docs/stages/01-getting-started.md) first and then use
the stage page to jump into the current source sections.

### Experienced Go developer

Jump to the stage you want to strengthen:

- concurrency: [Stage 07: Concurrency](./docs/stages/07-concurrency.md)
- quality and test: [Stage 08: Quality & Test](./docs/stages/08-quality-test.md)
- architecture: [Stage 09: Architecture](./docs/stages/09-architecture.md)
- production: [Stage 10: Production](./docs/stages/10-production.md)

## What You Will Build

You will work through:

- beginner-friendly exercises and starter projects
- parsers, filesystem tools, and CLI utilities
- HTTP services and database-backed applications
- concurrency pipelines, worker pools, and timeout-aware clients
- profiling, testing, and benchmark-driven improvements
- structured logging, graceful shutdown, gRPC, and deployment-ready workflows

## Current Docs

These docs are still useful:

| Document | Purpose |
| --- | --- |
| [LEARNING-PATH.md](./LEARNING-PATH.md) | current learning guide |
| [docs/stages/README.md](./docs/stages/README.md) | stage entry index |
| [docs/beta-public-architecture.md](./docs/beta-public-architecture.md) | explanation of source vs public architecture |
| [docs/curriculum/README.md](./docs/curriculum/README.md) | section-by-section curriculum map |
| [COMMON-MISTAKES.md](./COMMON-MISTAKES.md) | common Go bugs and fixes |
| [ROADMAP.md](./ROADMAP.md) | public roadmap |
| [CHANGELOG.md](./CHANGELOG.md) | change history |
| [CONTRIBUTING.md](./CONTRIBUTING.md) | contribution workflow |
| [RELEASE.md](./RELEASE.md) | release process |

Maintainers and contributors can follow planning on
[`planning/v2`](https://github.com/rasel9t6/the-go-engineer/tree/planning/v2/docs/v2).

## Current Repo Reality

- the 11-stage model is the public navigation truth
- the 11-stage model is the learner-facing navigation truth
- the root stage folders from `01-getting-started` through `11-flagship` are now the canonical
  public source surfaces
- some older top-level folders remain only as legacy redirects or contributor history

For the full explanation of that relationship, see
[docs/beta-public-architecture.md](./docs/beta-public-architecture.md).

## Run Lessons And Exercises

```bash
# run a lesson
go run ./01-getting-started/2-hello-world

# run a starter exercise
go run ./03-functions-errors/7-order-summary/_starter

# compare with the canonical solution
go run ./03-functions-errors/7-order-summary
```

## Validate The Repo

```bash
go run ./scripts/validate_curriculum.go
go test ./...
```

## Windows Note

Some database and capstone paths use `go-sqlite3`, which requires CGO and a C compiler.
If you are on Windows, WSL2 is the smoothest setup for those paths.

## License

This project is licensed under the **The Go Engineer License (TGE License) v1.0**.

- Free for personal, educational, and non-commercial use
- Commercial use requires permission
