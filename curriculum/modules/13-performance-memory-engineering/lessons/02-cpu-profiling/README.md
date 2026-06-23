# CPU profiling

## Mission

Understand and apply CPU profiling in the context of professional Go software engineering.

## Prerequisites

- core-13-01

## Mental Model

CPU profiling is like a photographer taking a flash photo of the CPU every 10ms. Each photo captures 'what function is the CPU executing right now?' After 3,000 photos (30 seconds), the stack trace that appears in 37% of photos likely consumes ~37% of CPU. It is not exact — sampling error exists — but with enough samples, the error is small enough to guide optimization decisions. Functions that never appear in any photo consume negligible CPU and should not be optimized first.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

When pprof.StartCPUProfile is called, the runtime calls runtime.SetCPUProfileRate(100) which programs the OS interval timer (setitimer on Linux, ITIMER_PROF) to deliver SIGPROF at 100Hz. The signal handler (sigprof in runtime/signal_unix.go) reads the interrupted thread's m (OS thread) and gp (goroutine pointer) from thread-local storage. It then calls gentraceback to unwind the goroutine's stack: starting from the PC/SP saved in the signal context, it walks frame-by-frame using Go's stack map metadata. Each frame's PC is mapped to a function name + line number via the gopclntab (Go program counter line number table). The resulting stack is hashed with a FNV-1a hash into a fixed-size hash table. If the bucket exists, the count is incremented; if not, a new bucket is allocated. On StopCPUProfile, the runtime serializes all buckets into a protobuf-format profile (pprof's proto format defined in src/runtime/pprof/proto.go).

## Run Instructions

```bash
go run ./curriculum/modules/13-performance-memory-engineering/lessons/02-cpu-profiling
go test ./curriculum/modules/13-performance-memory-engineering/lessons/02-cpu-profiling
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Profiling a cold process — a CPU profile collected when the service just started and has no request traffic shows runtime startup code (GC initialization, map zeroing, type registration) that disappears after warmup. Always profile under realistic load after the process has handled at least a few thousand requests.
- Optimizing functions that consume <5% of CPU — optimizing a function that appears at 3% in the profile improves overall performance by at most 3%, but the optimization itself may introduce bugs or complexity. Focus on the top 3-5 functions that cumulatively account for >50% of samples.
- Misreading flat vs cumulative time — flat time is samples where the function was on top of the stack (it was executing its own code). Cumulative time includes samples where the function was anywhere on the stack (it called something that consumed CPU). A high flat time means the function is CPU-intensive itself. A low flat but high cumulative time means the function calls something expensive.

## In Production

CPU profiling is the standard first step in Go performance investigations. Every major Go shop — Uber, Datadog, Cloudflare, Google — has an internal runbook that starts with 'collect a CPU profile.' Kubernetes sidecars like the Google Cloud Profiler agent periodically pull pprof endpoints. In incident response, a CPU profile is the fastest way to determine whether a latency spike is caused by CPU contention or I/O blocking.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-13-03`.
