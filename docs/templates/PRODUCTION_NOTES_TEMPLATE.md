# Production Notes Template

## Purpose

Every lesson should include a "In Production" section that teaches what can go wrong in real-world scenarios. This is what differentiates a syntax tutorial from an engineering curriculum.

---

## Template Structure

```markdown
## ⚠️ In Production

### This Can Cause...

- [Issue 1]
- [Issue 2]
- [Issue 3]

### Avoid Using This In...

- [Scenario 1] → Use [alternative] instead
- [Scenario 2] → Use [alternative] instead

### Common Production Bugs

- Bug 1: [Description] → Fix: [How to avoid]
- Bug 2: [Description] → Fix: [How to avoid]

### Monitoring Suggestions

- Track [metric name] for this operation
- Alert if [threshold]
```

---

## Example: Slices

````markdown
## ⚠️ In Production

### This Can Cause...

- **Memory leak**: If you keep appending to a slice and never clear it,
  the underlying array grows indefinitely. The slice header might be reset,
  but the underlying array is still referenced.
- **Unbounded growth**: In a hot path, append() without capacity
  pre-allocation causes multiple reallocations.
- **Shared backing array**: Sub-slices share memory with original slice.
  Modifying one can unexpectedly modify another.

### Avoid Using This In...

- High-frequency loops → Pre-allocate with make()
- Long-running services → Reset slice periodically or use circular buffer
- Concurrency → Use proper synchronization

### Common Production Bugs

1. **Slice leak**:
   ```go
   func process() {
       // Bug: slice keeps growing
       var items []Item
       for {
           items = append(items, getItem())
       }
   }
   ```
````

Fix: Use fixed buffer or reset periodically.

2. **Memory bloat**:
   ```go
   func getItems() []Item {
       result := make([]Item, 0, 1000)  // Pre-allocate
       // ...
   }
   ```

### Monitoring Suggestions

- Monitor slice length growth over time
- Alert if any slice grows beyond expected max
- Track memory allocation rate

````

---

## Example: HTTP Server

```markdown
## ⚠️ In Production

### This Can Cause...
- **Memory exhaustion**: Without timeouts, slow clients can hold
  connections indefinitely, eventually exhausting resources.
- **Goroutine leak**: If handlers spawn goroutines that never complete
  (no timeout), you get goroutine leak.
- **File descriptor exhaustion**: Each connection uses a file descriptor.
  Without limits, system runs out.

### Avoid Using This In...
- Production APIs → Always set ReadTimeout, WriteTimeout, IdleTimeout
- Public endpoints → Rate limiting required
- Long-running requests → Context with timeout

### Common Production Bugs
1. **No timeouts**:
   ```go
   // BAD: No timeout
   http.ListenAndServe(":8080", mux)

   // GOOD: With timeouts
   srv := &http.Server{
       Addr:         ":8080",
       Handler:      mux,
       ReadTimeout:  5 * time.Second,
       WriteTimeout: 10 * time.Second,
       IdleTimeout:  120 * time.Second,
   }
````

2. **Slowloris attack**: Client sends headers slowly to keep connection open.
   Fix: ReadHeaderTimeout

### Monitoring Suggestions

- Track active connections
- Track goroutine count
- Track request duration p99
- Alert on connection saturation

````

---

## Example: Concurrency

```markdown
## ⚠️ In Production

### This Can Cause...
- **Goroutine leak**: Forgetting to signal goroutine completion
  (e.g., not closing channel, not using WaitGroup properly).
- **Race conditions**: Data races cause intermittent, hard-to-debug bugs.
- **Deadlock**: Circular dependency between channels or mutexes.

### Avoid Using This In...
- Unbounded goroutines → Always use worker pool / semaphore
- Sharing state between goroutines → Use channels or sync.Mutex
- Infinite channels → Always close or use buffered with backpressure

### Common Production Bugs
1. **Goroutine leak**:
   ```go
   // BAD: No way to stop goroutine
   go func() {
       for {
           doWork()
       }
   }()

   // GOOD: With context cancellation
   go func() {
       for {
           select {
           case <-ctx.Done():
               return
           default:
               doWork()
           }
       }
   }()
````

2. **Channel deadlock**:
   ```go
   // BAD: Sending to unbuffered channel with no receiver
   ch := make(chan int)
   ch <- 1  // Blocks forever if no receiver!
   ```

### Monitoring Suggestions

- Monitor goroutine count over time
- Use race detector in tests
- Track channel buffer saturation
- Log goroutine start/stop for debugging

````

---

## Example: Error Handling

```markdown
## ⚠️ In Production

### This Can Cause...
- **Error loss**: Wrapping errors without %w loses the original error chain.
- **Silent failures**: Ignoring errors with _ means failures go unnoticed.
- **Information leak**: Returning raw database errors exposes internals.

### Avoid Using This In...
- API responses → Never return raw internal errors
- Production code → Never use _ to ignore errors
- Logging → Don't log same error multiple times

### Common Production Bugs
1. **Silent failure**:
   ```go
   // BAD: Error ignored
   result, _ := someFunction()

   // GOOD: At minimum, log it
   result, err := someFunction()
   if err != nil {
       log.Printf("warning: %v", err)
   }
````

2. **Error information leak**:

   ```go
   // BAD: Exposes internal details
   return err

   // GOOD: Generic error for external
   return fmt.Errorf("internal server error")
   ```

### Monitoring Suggestions

- Track error rate by type
- Alert on spike in error rate
- Log error chain for debugging
- Track specific error codes

```

---

## Section Checklist

Every section should answer:

- [ ] What can go wrong in production?
- [ ] What should NOT be used in which scenarios?
- [ ] What are common bugs?
- [ ] What should be monitored?

---

*This template should be added to every lesson in the curriculum.*
```
