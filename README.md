# The Go Engineer

[![CI](https://github.com/rasel9t6/the-go-engineer/actions/workflows/ci.yml/badge.svg)](https://github.com/rasel9t6/the-go-engineer/actions)
[![License: TGE License v1.0 (Non-Commercial)](https://img.shields.io/badge/License-TGE_v1.0-red.svg?style=for-the-badge&logo=data:image/svg%2Bxml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0id2hpdGUiPjxwYXRoIGQ9Ik0xNCAySDZhMiAyIDAgMCAwLTIgMnYxNmEyIDIgMCAwIDAgMiAyaDEyYTIgMiAwIDAgMCAyLTJWOGwtNi02em0tMSAxLjVMMTguNSA5SDEzVjMuNXpNNiAyMFY0aDV2Nmg2djEwSDZ6bTItOGg4djJIOHYtMnptMCA0aDh2Mkg4di0yeiIgZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiLz48L3N2Zz4%3D)](#license)
[![GitHub Sponsors](https://img.shields.io/badge/sponsor-30363D?style=for-the-badge&logo=GitHub-Sponsors&logoColor=%23EA4AAA)](https://github.com/sponsors/rasel9t6)
[![Patreon](https://img.shields.io/badge/Patreon-F96854?style=for-the-badge&logo=patreon&logoColor=white)](https://patreon.com/rasel9t6)

Learn Go by building, testing, and operating real software.

The Go Engineer is a repo-first Go software engineering learning system built around a progressive 5-Phase, 21-Section architecture. 

It takes you from absolute beginner syntax to building, testing, and deploying a production-grade SaaS Backend ("GoScale").

For a full breakdown of the curriculum architecture and roadmap, read [ARCHITECTURE.md](./ARCHITECTURE.md).

## Start Here

Pick the release channel that matches what you want:

- `release/v1`: the stable v1 line for learners who want the current supported experience
- `release/v2`: the current public beta checkpoint for the v2 line
- `main`: the active beta implementation branch and the fastest-moving line

## Quick Start

```bash
git clone https://github.com/rasel9t6/the-go-engineer.git
cd the-go-engineer
go mod download
go version
go run ./01-foundations/
```

## Curriculum Overview (5 Phases)

The curriculum is organized into 5 progressive phases spanning 21 sections. See [ARCHITECTURE.md](./ARCHITECTURE.md) for the exact section breakdowns.

1. **Phase 1: Syntax Foundation** - Compiling, types, control flow, functions, generics.
2. **Phase 2: Engineering Foundation** - Structs, methods, interfaces, errors, modularity.
3. **Phase 3: Production Engineering** - Concurrency, testing, context, IO.
4. **Phase 4: Systems Engineering** - Backend, APIs, gRPC, security, deployment.
5. **Phase 5: Flagship Project** - GoScale, the modular monolith capstone.

## What You Will Build

You will work through:

- beginner-friendly exercises and starter projects
- parsers, filesystem tools, and CLI utilities
- HTTP services and database-backed applications
- concurrency pipelines, worker pools, and timeout-aware clients
- profiling, testing, and benchmark-driven improvements
- structured logging, graceful shutdown, gRPC, and deployment-ready workflows
- **GoScale**: A full enterprise SaaS backend.

## Current Docs

These docs guide you through the learning system:

| Document | Purpose |
| --- | --- |
| [ARCHITECTURE.md](./ARCHITECTURE.md) | curriculum blueprint and vision |
| [LEARNING-PATH.md](./LEARNING-PATH.md) | current learning guide and entry points |
| [CURRICULUM-BLUEPRINT.md](./CURRICULUM-BLUEPRINT.md) | teaching and lesson contract standards |
| [CODE-STANDARDS.md](./CODE-STANDARDS.md) | engineering standards enforced in the repo |
| [TESTING-STANDARDS.md](./TESTING-STANDARDS.md) | testing coverage and patterns |
| [COMMON-MISTAKES.md](./COMMON-MISTAKES.md) | common Go bugs and fixes |
| [ROADMAP.md](./ROADMAP.md) | public roadmap |
| [CONTRIBUTING.md](./CONTRIBUTING.md) | contribution workflow |

## Run Lessons And Exercises

```bash
# run a lesson
go run ./01-foundations/

# run tests
go test ./...
```

## Validate The Repo

```bash
go run ./scripts/validate_curriculum.go
make test
make lint
```

## Windows Note

Some database and capstone paths use `go-sqlite3`, which requires CGO and a C compiler.
If you are on Windows, WSL2 is the smoothest setup for those paths.

## License

This project is licensed under the **The Go Engineer License (TGE License) v1.0**.

- Free for personal, educational, and non-commercial use
- Commercial use requires permission
