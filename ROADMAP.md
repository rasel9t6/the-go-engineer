# Roadmap

This document tracks what is built, what is in progress, and what is planned.

## Status Legend

| Symbol | Meaning |
|--------|---------|
| ✅ | Complete and tested |
| 🚧 | In progress |
| 📋 | Planned |
| 💡 | Under consideration |

---

## Phase 0 — Setup ✅
- `00-getting-started`: installation, hello world, compilation model, dev tools

## Phase 1 — Fundamentals ✅
- `01-language-basics`: variables, types, zero values, constants, iota
- `02-control-flow`: for, if/else, switch, line-of-sight principle
- `03-collections-and-pointers`: arrays, slices, maps, pointers, escape analysis
- `04-functions-and-errors`: multi-return, closures, defer, panic/recover, error wrapping
- `07-strings-and-text`: string internals, regex, templates, Builder

## Phase 2 — Core Go ✅
- `05-types-and-interfaces`: structs, methods, interfaces, generics
- `06-composition-and-embedding`: composition vs inheritance
- `11-encoding`: JSON marshal/unmarshal, streaming, base64
- `15-time-and-scheduling`: time, formatting, timers, tickers

## Phase 3 — Advanced Patterns ✅
- `09-concurrency`: goroutines, WaitGroup, channels, select, race conditions, sync primitives
- `10-filesystem`: file I/O, paths, directories, embed, io.Reader/Writer patterns

## Phase 4 — Production Engineering ✅
- `08-modules-and-dependencies`: go.mod, versioning, replace
- `12-databases`: sql.DB, connection pooling, transactions, repository pattern
- `13-web-masterclass`: routing, DI, templates, middleware, sessions, auth, forms, CRUD, pagination, comments
- `14-testing`: unit tests, table-driven, HTTP handler tests, benchmarking
- `16-http-clients-and-mocking`: http.Client, dependency injection, testify/mock

## Phase 5 — Expert Patterns ✅
- `17-context`: Background, WithCancel, WithTimeout, WithValue
- `18-package-design`: naming, visibility, internal/, project layout
- `19-cli-tools`: os.Args, flag, subcommands, file organiser

## Phase 6 — The Production Engineer ✅
- `20-docker-and-deployment`: multi-stage Dockerfile, layer caching
- `21-database-migrations`: golang-migrate, embed, schema evolution
- `22-enterprise-capstone`: full REST API, PostgreSQL, Docker Compose

---

## Phase 7 — Senior Engineer Patterns 🚧

### `23-structured-logging` ✅
- [x] `1-slog-basics`: TextHandler, JSONHandler, levels, attrs, groups, With
- [x] `2-context-logger`: request-scoped logger, HTTP middleware pattern
- [x] `3-custom-handler`: implementing slog.Handler, MultiHandler
- [x] `4-zerolog-comparison`: allocation-free logging, when to choose zerolog

### `24-errgroup-and-pools` ✅
- [x] `1-errgroup`: errgroup.Group, first-error collection, errors.Join, TryGo
- [x] `2-errgroup-context`: WithContext, automatic cancellation, fan-out pipeline
- [x] `3-sync-pool`: Get/Put lifecycle, buffer pool, struct pool, benchmarks
- [x] `4-bounded-pipeline-exercise`: Image Resizer solution with errgroup.SetLimit
- [x] `5-url-checker-exercise`: exercise + solution

### `25-profiling` 🚧
- [x] `1-cpu-profile`: pprof.StartCPUProfile, go tool pprof, flame graphs
- [ ] `2-memory-profile`: heap vs allocs, inuse_objects, runtime.MemStats
- [x] `3-http-pprof`: net/http/pprof, two-port pattern, security
- [ ] `4-trace`: runtime/trace, goroutine scheduling, GC events, latency analysis

### `26-grpc` 🚧
- [x] `proto/`: order.proto service definition
- [x] `1-unary/server`: implementing Server interface, status codes, interceptors
- [x] `1-unary/client`: client stub, context timeout, metadata, error handling
- [ ] `2-streaming/server`: server-side streaming RPC
- [ ] `2-streaming/client`: consuming a streaming response
- [ ] `3-bidi/`: bidirectional streaming
- [ ] `4-interceptors/`: auth interceptor, rate limiting interceptor

### `27-graceful-shutdown` ✅
- [x] `1-signal-context`: signal.NotifyContext, SIGTERM, shutdown deadline
- [x] `2-http-server`: http.Server.Shutdown, ErrServerClosed, server timeouts
- [x] `3-capstone`: complete pattern: readiness → HTTP drain → workers → DB → exit

---

## Phase 8 — Planned 📋

### `28-go-generate`
- `//go:generate` directive, running tools as part of the build
- `mockery` for generating testify mocks from interfaces
- `stringer` for enum String() methods
- `sqlc` for generating type-safe database code from SQL

### `29-cloud-native` 💡
- OpenTelemetry traces and metrics
- Kubernetes health probes (liveness vs readiness)
- Graceful shutdown for gRPC servers
- Config management with environment variables and Viper

### `30-performance-patterns` 💡
- `atomic.Value` for lock-free config hot-reload
- `maps` and `slices` packages (Go 1.21+)
- `cmp` package for ordered comparisons (Go 1.21+)
- Range over functions (Go 1.23+)
- SIMD-friendly data layout patterns

---

## Windows Setup Note

The `12-databases` and `22-enterprise-capstone` sections depend on `github.com/mattn/go-sqlite3` which requires CGO and a C compiler. On Windows without WSL:

1. Install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) or [MinGW-w64](https://www.mingw-w64.org/)
2. Ensure `gcc` is in your PATH: `gcc --version`
3. Set `CGO_ENABLED=1` in your environment

Alternatively, use WSL2 (recommended). All other sections work on Windows without a C compiler.

---

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) for how to add new lessons, exercises, and sections.

To propose a new section or report a curriculum gap, open an issue with the `[PROPOSAL]` label.
