# The Go Engineer: Learn Go by Building Real Projects

[![CI](https://github.com/rasel9t6/the-go-engineer/actions/workflows/ci.yml/badge.svg)](https://github.com/rasel9t6/the-go-engineer/actions)
[![License: TGE License v1.0 (Non-Commercial)](https://img.shields.io/badge/License-TGE_v1.0-red.svg?style=for-the-badge&logo=data:image/svg%2Bxml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCIgZmlsbD0id2hpdGUiPjxwYXRoIGQ9Ik0xNCAySDZhMiAyIDAgMCAwLTIgMnYxNmEyIDIgMCAwIDAgMiAyaDEyYTIgMiAwIDAgMCAyLTJWOGwtNi02em0tMSAxLjVMMTguNSA5SDEzVjMuNXpNNiAyMFY0aDV2Nmg2djEwSDZ6bTItOGg4djJIOHYtMnptMCA0aDh2Mkg4di0yeiIgZmlsbC1ydWxlPSJldmVub2RkIiBjbGlwLXJ1bGU9ImV2ZW5vZGQiLz48L3N2Zz4%3D)](#-license)
[![GitHub Sponsors](https://img.shields.io/badge/sponsor-30363D?style=for-the-badge&logo=GitHub-Sponsors&logoColor=#EA4AAA)](https://github.com/sponsors/rasel9t6)
[![Patreon](https://img.shields.io/badge/Patreon-F96854?style=for-the-badge&logo=patreon&logoColor=white)](https://patreon.com/rasel9t6)

Welcome to **The Go Engineer** — the definitive open-source Go curriculum. Every section teaches through **practical examples, real-world projects, and hands-on exercises** — not just syntax. You'll build servers, CLI tools, concurrent pipelines, REST APIs, and production-grade applications while learning the engineering depth behind every concept.

## Quick Start

```bash
# 1. Install Go: https://go.dev/dl/  (see 00-getting-started for detailed instructions)
# 2. Clone this repository
git clone https://github.com/rasel9t6/the-go-engineer.git
cd the-go-engineer

# 3. Download dependencies
go mod download

# 4. Verify Go is working
go version

# 5. Run your first program
go run ./00-getting-started/2-hello-world
```

## Reference Documents

| Document | Purpose |
| -------- | ------- |
| [COMMON-MISTAKES.md](./COMMON-MISTAKES.md) | 15 most common Go bugs with fixes and section cross-references |
| [ROADMAP.md](./ROADMAP.md) | What is built, in progress, and planned |
| [CHANGELOG.md](./CHANGELOG.md) | History of additions and bug fixes |
| [CONTRIBUTING.md](./CONTRIBUTING.md) | How to add lessons, exercises, and sections |

## Who is This For?

- **Complete beginners** — Never programmed before? Start at Section 00. Every line is explained.
- **Developers from other languages** — Know Python/JS/Java? Start at Section 01. Skim the basics, deep-dive into Go-specific patterns.
- **Go developers leveling up** — Already write Go? Jump to Sections 09+ for concurrency, testing, and production patterns.

## The Structured 4-Phase Learning Path

This repository follows a strict **Beginner → Expert** progression. Every section builds on the previous one.

### Phase 0: Setup

Get your development environment ready.

- `00-getting-started`: Installation, Hello World, how Go works, dev environment setup

### Phase 1: Fundamentals

Start here to understand the syntax, memory model, and error handling philosophies of Go.

- `01-language-basics`: Variables, types, scope, zero values, and formatting
- `02-control-flow`: Loops, if/switch statements, branching
- `03-collections-and-pointers`: Arrays, Slices, Maps, and Pointers (Stack vs Heap)
- `04-functions-and-errors`: First-class functions, closures, defer mechanics, and idiomatic error handling
- `07-strings-and-text`: String internals (UTF-8 bytes), regex, and buffers

### Phase 2: Core Go (Architecture & Types)

Master Go's approach to Object-Oriented Programming and domain modeling.

- `05-types-and-interfaces`: Custom types, structs, methods (pointer vs value receivers), duck typing
- `06-composition-and-embedding`: Composition over inheritance
- `11-encoding`: JSON serialization/deserialization mechanics
- `15-time-and-scheduling`: `time`, `context` trees, and tickers

### Phase 3: Advanced Patterns

Dive into Go's superpower: Concurrency and low-level I/O.

- `09-concurrency`: Goroutines, channel blocking, WaitGroups, `select` deep-dives, sync primitives (`Mutex`, `Map`)
- `10-filesystem`: Low-level I/O, `io.Reader`/`io.Writer` patterns, and the `embed` package

### Phase 4: Production Engineering

Learn how to build, test, and deploy Google-grade applications.

- `08-modules-and-dependencies`: `go.mod`, dependency management, versioning
- `12-databases`: `database/sql`, connection pooling, SQLite, Repository pattern
- `13-web-masterclass`: Comprehensive Web Development (Routing → Auth → Full App)
- `14-testing`: Unit testing, structured benchmarks (`testing.B`), memory allocation profiling
- `16-http-clients-and-mocking`: Calling APIs, dependency injection for testing, `testify/mock`

### Phase 5: Expert Patterns

Master the patterns that distinguish senior Go engineers.

- `17-context`: `context.Context` deep-dive — cancellation, timeouts, request-scoped values
- `18-package-design`: Naming conventions, visibility, `internal/`, standard project layout
- `19-cli-tools`: Building command-line tools with `flag`, subcommands, and exit codes

### Phase 7: Senior Engineer Patterns

- `23-structured-logging`: `slog` basics, context-keyed logger, custom `slog.Handler`, zerolog comparison
- `24-errgroup-and-pools`: `errgroup.Group`, errgroup+context pipelines, `sync.Pool` for zero-allocation buffering
- `25-profiling`: CPU profiles, memory profiles, live `net/http/pprof` endpoint, `go tool pprof` workflow
- `26-grpc`: proto3 service definition, unary and streaming RPC, interceptors (middleware for gRPC)
- `27-graceful-shutdown`: `signal.NotifyContext`, `http.Server.Shutdown`, complete zero-downtime Kubernetes shutdown

### Phase 6: The Production Engineer

Master the tools required to deploy scalable Go services to the cloud.

- `20-docker-and-deployment`: Multi-stage Dockerfiles, layer caching, `docker-compose` orchestration
- `21-database-migrations`: Schema evolution using `golang-migrate` and `//go:embed` SQL embedding
- `22-enterprise-capstone`: The ultimate multi-package REST API (PostgreSQL + Middleware + DB Migrations + Docker)

## Projects & Exercises

Each section culminates in a hands-on project to test your understanding:

| Section | Exercise | Description |
| ------- | -------- | ----------- |
| **01** Language Basics | `4-application-logger` | Application Logger with severity levels |
| **02** Control Flow | `4-pricing-calculator` | Pricing Calculator engine |
| **03** Collections | `6-contact-manager` | Slice-based Contact Manager System |
| **04** Functions & Errors | `8-error-handling` | Custom mathematical error handling |
| **05** Types & Interfaces | `6-payroll-processor` | Polymorphic User Payroll Processor |
| **06** Composition | `3-bank-account` | Bank Account System with deposits/withdrawals |
| **07** Strings & Text | `6-log-parser` | Log File Parsing System |
| **09** Concurrency | `7-downloader` | Concurrent Multi-File Downloader |
| **10** Filesystem | `7-log-search` | Directory traversal log search tool |
| **11** Encoding | `6-config-parser` | JSON config file parser with validation |
| **12** Databases | `6-repository` | CRUD SQLite App using Repository Pattern |
| **13** Web Masterclass | `1-routing/exercise` | Multi-route Bookstore Web API |
| **15** Time & Scheduling | `7-reminder` | Console reminder with countdown timer |
| **16** HTTP Clients | `6-testify-mock` | Mocking an external REST API Data Fetcher |
| **17** Context | `5-timeout-client` | Timeout-aware HTTP API client |
| **19** CLI Tools | `4-file-organizer` | CLI file organizer by extension |
| **22** The Capstone | `cmd/api` | **The Multi-Package Docker Enterprise Backend** |
| **23** Structured Logging | `2-context-logger` | HTTP middleware that injects request-scoped logger into context |
| **24** errgroup & sync.Pool | `4-bounded-pipeline-exercise` | Image resizer with bounded concurrency using `errgroup.SetLimit` |
| **24** errgroup & sync.Pool | `5-url-checker-exercise` | Concurrent URL health checker with bounded concurrency and pooled clients |
| **25** Profiling | `1-cpu-profile` | Profile slow vs fast log processor, read flame graph in go tool pprof |
| **26** gRPC | `1-unary` | Type-safe OrderService client + server with status codes and interceptors |
| **27** Graceful Shutdown | `3-capstone` | Complete shutdown: readiness probe → HTTP drain → worker drain → DB close |

## How to Use This Repository

The best way to learn is by **reading the inline comments** and **running the code**.

```bash
# Run any lesson
go run ./SECTION/LESSON

# Examples:
go run ./00-getting-started/2-hello-world
go run ./01-language-basics/1-variables
go run ./09-concurrency/3-channels
go run ./13-web-masterclass/1-routing

# Section 23 — Structured Logging
go run ./23-structured-logging/1-slog-basics
go run ./23-structured-logging/2-context-logger    # then: curl http://localhost:8080/api/orders/42

# Section 24 — errgroup & sync.Pool
go run ./24-errgroup-and-pools/1-errgroup
go run ./24-errgroup-and-pools/5-url-checker-exercise  # Try the exercise
go run ./24-errgroup-and-pools/5-url-checker-exercise/_starter  # See the solution

# Section 25 — Profiling
go run ./25-profiling/1-cpu-profile               # then: go tool pprof -http=:8090 cpu.prof
go run ./25-profiling/3-http-pprof                # then: go tool pprof http://localhost:6060/debug/pprof/profile?seconds=5

# Section 26 — gRPC (two terminals)
go run ./26-grpc/1-unary/server                   # Terminal 1
go run ./26-grpc/1-unary/client                   # Terminal 2

# Section 27 — Graceful Shutdown
go run ./27-graceful-shutdown/1-signal-context     # then: Ctrl+C to test
go run ./27-graceful-shutdown/2-http-server        # then: curl http://localhost:8080/api/slow & Ctrl+C
go run ./27-graceful-shutdown/3-capstone           # complete production pattern
```

### Self-Challenge Mode

Most exercises include a `_starter/` directory with TODO stubs:

```bash
# Try the exercise yourself first:
go run ./02-control-flow/4-pricing-calculator/_starter

# Then compare with the solution:
go run ./02-control-flow/4-pricing-calculator
```

For the grand finale, boot the entire Enterprise Backend cluster (Database + Migrations + API) using Docker:

```bash
# Run the massive Phase 6 Capstone project
cd 22-enterprise-capstone
docker-compose up -d --build
```

## Running the Tests

To verify your environment is set up correctly, run the full test suite:

```bash
# Run all tests
go test ./...

# Run tests with race detection
go test -race ./...

# Run benchmarks
go test -bench=. -benchmem -count=1 ./14-testing/benchmarks/
```

## Windows Users — CGO Note

Sections `12-databases` and `22-enterprise-capstone` use `github.com/mattn/go-sqlite3`,
which requires CGO and a C compiler. On Windows without WSL:

1. Install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/)
2. Verify: `gcc --version`
3. Set environment: `$env:CGO_ENABLED = "1"` (PowerShell)

All other sections (00–11, 13–27) work on Windows without a C compiler.
We recommend [WSL2](https://docs.microsoft.com/en-us/windows/wsl/) for the best experience.

## Requirements

- **Go 1.26+** (check with `go version`)
- **Git** (to clone the repository)
- A text editor (VSCode with Go extension recommended — see Section 00)

## 📜 License

This project is licensed under the **The Go Engineer License (TGE License) v1.0**.

- ✅ Free for personal, educational, and non-commercial use
- ❌ Commercial use is strictly prohibited without permission

### 🚫 Commercial Use Includes

- Paid courses, bootcamps, or training programs
- SaaS or product development
- Company/internal usage
- Freelance or client work
- Selling this project or derivatives

### 📧 Commercial License

To use this project commercially, contact:

[raselhossen052@gmail.com](mailto:raselhossen052@gmail.com)
