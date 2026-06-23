# Caching basics

## Mission

Understand and apply Caching basics in the context of professional Go software engineering.

## Prerequisites

- core-13-06

## Mental Model

A cache is a fast temporary storage that sits between the application and a slow data source. The drawer has limited space — when it fills up, the least-used tool goes back to the toolbox (LRU eviction). The key insight: caching only helps if the access pattern has locality — the same keys are requested repeatedly within the eviction window.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A Go map used as a cache is a hash table with O(1) average lookup, but it is not goroutine-safe without external synchronization. sync.Map uses a combination of atomic loads for read-mostly workloads and a mutex-protected 'dirty' map for writes. hashicorp/golang-lru implements LRU with a doubly-linked list + map: on access, the entry is moved to the front of the list; on eviction, the entry at the back is removed. TTL expiration in most implementations uses lazy eviction (check on Get) plus a periodic cleanup goroutine. The tradeoff is: lazy eviction is cheaper per-op but can leave stale entries in memory longer; periodic cleanup removes stale entries in batches but adds a mutex contention spike every N seconds. For very large caches, use a sharded design: divide the key space across N independent caches, each with its own lock, reducing lock contention by Nx.

## Run Instructions

```bash
go run ./curriculum/modules/13-performance-memory-engineering/lessons/07-caching-basics
go test ./curriculum/modules/13-performance-memory-engineering/lessons/07-caching-basics
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using sync.Map for read-heavy workloads with simple keys — sync.Map is optimized for two patterns: (1) append-only writes with many readers, (2) write-once/read-many. For a standard read-heavy workload (90% reads, single writer), a regular map with sync.RWMutex is 2-5x faster because sync.Map's atomic Load overhead outweighs mutex contention at low concurrency.
- Adding a cache without a TTL or eviction policy — an unbounded cache grows until the process OOMs. Every cache must have: max size (item count or byte limit), eviction policy (LRU, LFU, TTL), and a plan for stale data. A cache with no eviction policy is a memory leak waiting to happen.
- Caching at the wrong layer — caching a computed value instead of the expensive input-to-output mapping. If the expensive work is json.Unmarshal of the same response body, cache the parsed result (Go struct), not the raw HTTP response. Caching at the right layer maximizes the cost saved per cache hit.

## In Production

In-memory caching is the most common optimization technique in Go production services. Standard patterns: (1) per-request cache (sync.Once or map inside a request context), (2) per-process cache (sync.Map or golang-lru shared by all requests), (3) distributed cache (Redis with go-redis client). Kubernetes uses in-memory caching for API object deserialization. The Go playground caches compiled binaries. Every Go HTTP service with a database should have some form of in-memory cache.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-13-08`.
