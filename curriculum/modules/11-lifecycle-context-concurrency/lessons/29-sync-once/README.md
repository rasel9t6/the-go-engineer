# sync.Once

## Mission

Understand and apply sync.Once in the context of professional Go software engineering.

## Prerequisites

- core-11-28

## Mental Model

sync.Once is a one-shot gate that guarantees a function executes exactly once, no matter how many goroutines call Do simultaneously. After the function completes, all subsequent Do calls are no-ops.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The internal implementation uses a 32-bit counter with atomic operations. The fast path is a single atomic.Load on the done flag. When done is 0, Do acquires a mutex, double-checks the flag, and executes f. After f returns, done is atomically set to 1 and the mutex is released. This is the textbook double-checked locking pattern, made safe by Go's memory model.

## Run Instructions

```bash
go run ./curriculum/modules/11-lifecycle-context-concurrency/lessons/29-sync-once
go test ./curriculum/modules/11-lifecycle-context-concurrency/lessons/29-sync-once
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using sync.Once inside a function that is called in a loop — the Do method only runs the first invocation, so subsequent calls are silently skipped.
- Wrapping sync.Once.Do(f) in a mutex because the developer assumes Do itself needs external synchronization — the whole point of sync.Once is that Do is goroutine-safe.
- Calling sync.Once.Do(f) where f panics — the panic propagates to the caller, but the underlying Once still marks the operation as done, so recovery and retry are impossible.

## In Production

sync.Once is used for lazy initialization of database/sql connection pools, TLS certificate caches, configuration loading in CLI tools, one-time registration of custom metrics in Prometheus, and lazy computation in caching layers.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `module checkpoint`.
