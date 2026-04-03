# Section 24: errgroup & sync.Pool

## Beginner → Expert Mapping

| Topic | Level | Importance | Engineering Concept |
|-------|-------|------------|---------------------|
| `errgroup.Group` | Intermediate | **Critical** | Concurrent error collection, idiomatic WaitGroup replacement |
| `errgroup` + context | Advanced | **Critical** | Automatic cancellation on first error |
| `sync.Pool` | Advanced | High | Object reuse, GC pressure reduction |
| Bounded worker pool | Expert | High | Semaphore + errgroup, back-pressure |

## Engineering Depth

`errgroup` is the missing piece between WaitGroup and channels. WaitGroup tells you *when* goroutines finish but discards their errors. Channels collect errors but require manual synchronization. `errgroup.Group` does both: it waits for all goroutines and returns the first non-nil error, cancelling all remaining work via context.

`sync.Pool` is the most impactful memory optimization available in Go. A pool holds recycled objects. Instead of `make([]byte, 4096)` on every request (triggering GC), you `Get()` a pre-allocated buffer, use it, and `Put()` it back. The standard library uses pools everywhere: `fmt`, `encoding/json`, `net/http` all have internal pools.

**When to use sync.Pool:**
- Object allocation appears in pprof heap profile hot path
- Objects are temporary: used briefly, then discarded
- Object size is non-trivial: byte buffers, structs with multiple fields

**When NOT to use sync.Pool:**
- Objects hold state between requests (race condition)
- Allocation is not the bottleneck (profile first!)
- Objects outlive the goroutine that created them

## Contents

| Directory | Topic | Level |
|-----------|-------|-------|
| `1-errgroup/` | errgroup replacement for WaitGroup | Intermediate |
| `2-errgroup-context/` | Automatic cancellation via Context | Advanced |
| `3-sync-pool/` | Reducing GC pressure with sync.Pool | Advanced |
| `4-exercise/` | Build a concurrent batch processor | Intermediate |

## How to Run

```bash
go run ./24-errgroup-and-pools/1-errgroup
go run ./24-errgroup-and-pools/2-errgroup-context
go run ./24-errgroup-and-pools/3-sync-pool
go test -bench=. -benchmem ./24-errgroup-and-pools/3-sync-pool
go run ./24-errgroup-and-pools/4-exercise/_starter
```

## References

- [golang.org/x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)
- [sync.Pool](https://pkg.go.dev/sync#Pool)
- [Go Blog: Profiling Go Programs](https://go.dev/blog/pprof)
