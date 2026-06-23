# Memory profiling

## Mission

Understand and apply Memory profiling in the context of professional Go software engineering.

## Prerequisites

- core-13-02

## Mental Model

A heap profile is a photograph of the heap at a single moment, showing 'what code allocated the memory that is still alive?' Each sample represents ~512KB of live allocation. If a function appears in 10 samples, it is responsible for ~5MB of live heap. The profile is a probability-weighted estimate: functions that allocate rarely may not appear in any sample, and functions that allocate heavily appear in proportion to their allocation.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Heap profiling is implemented in runtime/mprof.go. Each M (OS thread) has a per-thread allocation counter. When mallocgc is called, it increments the counter by the allocation size. When the counter exceeds MemProfileRate / 2 (biased for randomness), the runtime generates a random number and decides whether to sample. On sample, it calls recordstacktrace which captures up to 32 PC values from the current goroutine's stack. The record is stored in a global slice of MemProfileRecord entries. Each record contains: AllocBytes (total bytes allocated at this site), FreeBytes (bytes freed), AllocObjects (count allocated), FreeObjects (count freed), and the stack trace. On WriteHeapProfile, the runtime iterates all records, computes inuse = AllocBytes - FreeBytes, and serializes as a pprof profile. The profile type is 'heap' and each sample type has two units: 'alloc_objects'/'alloc_space' and 'inuse_objects'/'inuse_space'.

## Run Instructions

```bash
go run ./curriculum/modules/13-performance-memory-engineering/lessons/03-memory-profiling
go test ./curriculum/modules/13-performance-memory-engineering/lessons/03-memory-profiling
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Reading alloc_space instead of inuse_space — alloc_space shows the total allocations since start (cumulative bytes allocated over the process lifetime). inuse_space shows currently live allocations (allocated minus freed). For finding memory leaks, inuse_space is the relevant metric. alloc_space is useful for understanding allocation rate and GC pressure.
- Profiling heap at the wrong point in the request lifecycle — a heap profile collected during peak request processing shows inflated inuse because many requests are in-flight. A profile collected during idle shows only the baseline. Compare both to distinguish per-request allocation from persistent allocation.
- Forgetting that heap profiles show allocation call sites, not current owners — a heap profile shows which function allocated the memory, but does not show which code path currently holds the reference. A buffer allocated at connection creation and held by a leaked goroutine appears as 'allocated by conn.setup' even though the goroutine is the root of the leak.

## In Production

Memory profiling is the standard tool for debugging Go memory leaks and high memory usage. Kubernetes operators use heap profiles to set memory limits (the profile shows the live heap size under load). Datadog's continuous profiler collects heap profiles from Go services every minute. In incident response, the first question about an OOM kill is: 'collect a heap profile and look at inuse_space by function.'

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-13-04`.
