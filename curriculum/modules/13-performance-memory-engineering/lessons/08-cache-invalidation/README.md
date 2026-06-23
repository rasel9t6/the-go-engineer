# Cache invalidation

## Mission

Understand and apply Cache invalidation in the context of professional Go software engineering.

## Prerequisites

- core-13-07

## Mental Model

Cache invalidation is the problem of keeping a copy (cache) consistent with the source of truth (DB). If the erasure is forgotten, the whiteboard shows old data. If the erasure happens before the ledger update, the whiteboard is empty even though the update failed. If two people erase and write simultaneously, the final whiteboard state may match neither intended update. The three mechanisms are: TTL (the whiteboard self-erases after a timeout), active invalidation (someone erases when the ledger changes), and write-through (someone updates the whiteboard as part of the ledger update). Each has tradeoffs.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Cache invalidation in Go typically uses one of three patterns. (1) Cache-aside with delete: the application reads from cache first; on cache miss, reads from DB and populates the cache. On writes, it updates the DB and deletes the cache entry (or updates it). This is the most common pattern because it handles concurrent writes safely — deleting is idempotent, and the next read fetches fresh data. (2) Write-through: the application updates both cache and DB atomically. Go has no built-in distributed transaction support, so this requires coordinating a two-phase commit between cache (Redis) and DB (PostgreSQL) — typically implemented with the outbox pattern: write to a DB table, a separate process reads the outbox and updates the cache. (3) Write-behind: the application writes to cache immediately and queues the DB write. If the process crashes between cache write and DB write, the data is lost. Go's channels and worker pools implement this pattern naturally. The key invariant: in cache-aside, the cache is always a snapshot of the DB at some point in time; the invalidation strategy determines the staleness window, not the correctness.

## Run Instructions

```bash
go run ./curriculum/modules/13-performance-memory-engineering/lessons/08-cache-invalidation
go test ./curriculum/modules/13-performance-memory-engineering/lessons/08-cache-invalidation
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Setting a long TTL as the only invalidation mechanism — a 30-minute TTL means updates are invisible for up to 30 minutes. Users see stale data and report bugs that are 'already fixed.' TTL-only caches are acceptable only for data that changes rarely or where staleness is acceptable (e.g., 'top 10 articles' updated hourly).
- Deleting cache entries before writing the new value to the database — if the cache entry is deleted and the DB write fails, the next request sees a cache miss and the old data is never served again until a new write succeeds. The cache is now empty for this key, and the DB may be hit harder. Fix: delete after the DB write succeeds, not before.
- Invalidating a cache entry that was never cached — calling cache.Delete(key) for every DB write is wasteful if the key was never cached. Track whether a key has been cached (e.g., a bloom filter or a set of cached keys) before sending invalidation events. Unnecessary invalidation is harmless for correctness but creates write amplification on the cache.
- Using cache-aside without a version check — in cache-aside (lazy population), the cache stores whatever the DB returns. If two concurrent requests race: request A updates the DB, request B reads old data and writes the old value to cache — the old value overwrites the new one. Fix: use a version number or timestamp and only write to cache if the version matches.

## In Production

Cache invalidation strategies are a core architectural decision in every production Go service with caching. Kubernetes uses watch-based invalidation: clients watch for resource changes and invalidate local caches. Uber's caching layer uses a write-through pattern with Redis and a version matrix. The Go playground uses TTL-only (10 minutes) because the data changes rarely and staleness is acceptable. In distributed systems, cache invalidation is typically solved with: (1) Redis pub/sub for invalidation events, (2) database triggers that emit NOTIFY on row changes (LISTEN/NOTIFY), or (3) a version vector attached to each cache entry.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `module checkpoint`.
