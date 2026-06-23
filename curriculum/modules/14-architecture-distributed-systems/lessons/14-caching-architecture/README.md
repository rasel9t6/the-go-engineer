# Caching architecture

## Mission

Understand and apply Caching architecture in the context of professional Go software engineering.

## Prerequisites

- core-14-13

## Mental Model

Cache is a fast, small memory that sits in front of a slow, large storage. Frequently used files are on your desk (in cache) — you grab them instantly. If a file is not on your desk, you walk to the filing cabinet (database query). If someone asks for the same file right after you put it on your desk, they get it instantly. The key insight: you must update BOTH the filing cabinet AND the desk when a file changes, or the desk has stale information.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's singleflight.Group (golang.org/x/sync/singleflight) coalesces concurrent function calls with the same key. When 100 goroutines call g.Do("user:42", fn), only one executes fn, and all 100 receive the same result. This is the canonical Go solution for cache stampede prevention. The implementation: singleflight uses a sync.Mutex + map[string]*call. The first call creates a *call entry with a channel. Subsequent calls find the existing entry and wait on the channel. When the first call completes, it sends the result on the channel, waking all waiters. The entry is then deleted from the map.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/14-caching-architecture
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/14-caching-architecture
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Caching everything with the same TTL — user profile data (changes rarely) and inventory counts (change every minute) should have very different TTLs. Using a 5-minute TTL for everything means user profiles are fetched from the database on every request (no cache benefit), and inventory counts serve stale data for 5 minutes.
- Not handling cache stampedes — when a cache key expires and 100 concurrent requests all miss the cache, they all hit the database simultaneously, amplifying load by 100x. Fix: use a mutex or 'probabilistic early expiration' — one request regenerates the cache, others wait for it.
- Using cache as the source of truth — if Redis goes down and all cached data is lost, the service must reconstruct it from the database. If the cache is treated as the source of truth and the DB is a backup, cache loss causes data loss. Fix: the database is the source of truth; the cache is a performance optimization.

## In Production

Caching is the most common performance optimization in production Go services. Redis is the standard distributed cache, used by: Twitter (user timelines cached in Redis), GitHub (repository data cached in Redis), Stack Overflow (page content cached in Redis). In-memory caching (sync.Map or go-cache) is used for service-local data that does not need cross-instance consistency: feature flags, configuration, rate limit counters.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-15`.
