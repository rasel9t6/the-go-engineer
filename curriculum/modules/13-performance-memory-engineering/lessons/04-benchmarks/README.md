# Benchmarks

## Mission

Understand and apply Benchmarks in the context of professional Go software engineering.

## Prerequisites

- core-13-03

## Mental Model

A benchmark is a stopwatch that runs a function many times and reports the average time per run. The framework finds the right number of repetitions to get a stable average — just like measuring a runner's lap time by timing 100 laps, not 1 lap. The key insight: benchmarks measure the same code path repeatedly to cancel out measurement noise.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

When go test -bench runs, the testing framework parses benchmark function names, creates a testing.B for each, and calls the benchmark function. The framework starts with b.N = 1 and runs the function once. It measures the wall-clock duration with time.Now(). If the duration is below the minimum benchmark time (default 1 second), it computes a scaling factor: target_ns / measured_ns * 1.1 (10% safety margin). It then runs the benchmark again with the estimated N. It repeats until either the runtime is above the minimum time or N stops growing (for very fast operations, it may iterate this calibration loop multiple times). Internally, testing.B calls runtime.ReadMemStats before and after the benchmark loop to compute B/op and allocs/op. The reported value is the delta divided by N. The framework also tracks the number of iterations completed in each sub-benchmark and reports the minimum time across runs (not the average) to reduce noise from GC or OS scheduling.

## Run Instructions

```bash
go run ./curriculum/modules/13-performance-memory-engineering/lessons/04-benchmarks
go test ./curriculum/modules/13-performance-memory-engineering/lessons/04-benchmarks
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Benchmarking with compiler optimizations disabled — running a benchmark with -gcflags='-N -l' (disable optimization and inlining) measures the unoptimized code path, which is never what runs in production. Always benchmark with default compiler flags. The compiler inlines and optimizes differently at different optimization levels.
- Not resetting the timer before the measured loop — if the benchmark function has setup code (file reads, allocation of test data), that setup time is included in the reported ns/op unless b.ResetTimer() is called before the loop. The setup happens once but its time skews the per-op measurement.
- Benchmarking I/O-bound work — the -bench framework measures CPU time per operation. If the benchmarked function makes an HTTP call or reads from disk, the reported ns/op includes I/O wait time, which is non-deterministic and varies with system load. Use go test -bench for CPU-bound code only; use go test -bench with -count and statistical analysis for I/O-bound code.

## In Production

Benchmarks are the standard tool for Go performance regression testing. The standard library runs benchmarks on every commit. Kubernetes benchmarks critical code paths (API serialization, RBAC evaluation) and requires benchmark evidence for performance-sensitive changes. In the Go community, benchmarks are the currency of optimization discussions — claims like 'this library is faster' must include a benchmark run with -benchmem.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-13-05`.
