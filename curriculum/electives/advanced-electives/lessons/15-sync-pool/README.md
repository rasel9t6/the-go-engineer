# sync.Pool

## Mission

Understand and apply sync.Pool in the context of professional Go software engineering.

## Prerequisites

- elective-14

## Mental Model

sync.Pool is a temporary object cache per-P (per-processor). Each goroutine accesses its local P's pool with minimal synchronization. Get retrieves an object from the pool (or allocates a new one if empty). Put returns an object to the pool for reuse. The pool is drained during GC — objects are freed, and subsequent Get() calls allocate fresh objects. The pool is NOT a cache with predictable retention; it is an allocation amortization mechanism.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

sync.Pool maintains per-P private and shared lists. Get() first checks the P's private slot (fastest, no lock), then the P's shared list (fast, per-P lock), then steals from other P's shared lists (slower, cross-P lock), and finally calls New() to allocate. Put() stores in the P's private slot if empty, otherwise appends to the shared list. During each GC cycle, the pool's local caches are drained: all objects are freed unless referenced elsewhere.

## Run Instructions

```bash
Read the lesson and complete the practice task.
No automated test is required for this lesson.
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming sync.Pool is a persistent cache — objects in the pool are removed during GC; use a regular cache (e.g., LRU) for persistent caching.
- Putting objects without resetting them — the next Get() caller receives an object with residual state; always reset (buf.Reset(), obj.Clear()) before use.
- Using sync.Pool for objects that hold external resources — the pool does not call Close or cleanup; objects holding file handles or network connections leak if not explicitly closed.
- Not calling Put in a defer — if the function panics or returns early, the object is not returned to the pool, increasing GC pressure.
- Using sync.Pool for large objects — large objects in the pool increase GC scan time; pool only objects that are expensive to allocate AND cheap to reset.

## In Production

sync.Pool is used in every Go production service for: bytes.Buffer pooling (JSON/protobuf serialization), message struct pooling (event publishing), SQL row scanning buffers, template rendering buffers, and any hot-path allocation-heavy operation. Opslane uses sync.Pool for order event marshaling buffers and HTTP response writer buffers.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-16`.
