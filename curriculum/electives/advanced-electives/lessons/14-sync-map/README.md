# sync.Map

## Mission

Understand and apply sync.Map in the context of professional Go software engineering.

## Prerequisites

- elective-13

## Mental Model

sync.Map has two internal maps: a read map (accessed atomically, no lock) and a dirty map (accessed with a mutex). Reads first check the read map — if the key exists, the read returns instantly with no lock. If the key is not in the read map, the read falls back to the dirty map with a lock. Writes always go to the dirty map with a lock. Periodically, the dirty map is promoted to become the new read map (amended=true).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The read map is a pointer to a readOnly struct that contains a map and an amended flag. Load() atomically loads the read map pointer and checks for the key. If found (and not deleted), return the value without locking. If not found and amended is true, acquire the mutex, promote the dirty map to read, and retry. Store() acquires the mutex, writes to the dirty map, and sets amended=true. Delete() marks the key as deleted in the read map (without removing it) and removes it from the dirty map under the mutex.

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

- Using sync.Map when a regular map with sync.Mutex is simpler — sync.Map has a complex API (Load, Store, LoadOrStore, Delete, Range) that is harder to use than a regular map with a mutex.
- Using sync.Map for write-heavy workloads — sync.Map's performance degrades when writes are frequent because every write requires a mutex lock and periodic map promotion.
- Using sync.Map for small maps — a regular map with a mutex performs better for maps with fewer than 10 entries because the atomic read path overhead is not justified.
- Not using Range for iteration — sync.Map does not support for range; use Range(f func(key, value any) bool) instead, and return false to stop iteration.
- Using sync.Map as a typed map — Load returns any, requiring type assertions. Consider a typed wrapper around sync.Map or a regular map with a mutex for type safety.

## In Production

sync.Map is used in Go for: connection pools (HTTP client, database), configuration caches (read-mostly, updated infrequently), feature flag registries (set at startup, read by every request), metrics registries (registered once, read periodically by the metrics exporter), and DNS or IP caches.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-15`.
