# pprof

## Mission

Understand and apply pprof in the context of professional Go software engineering.

## Prerequisites

- core-12-09

## Mental Model

pprof is the process equivalent of an x-ray machine taking snapshots at regular intervals. Each snapshot captures 'what line of code is the CPU executing right now?' (CPU profile) or 'what code path allocated this memory?' (heap profile). After enough snapshots (thousands), the distribution of stack traces accurately reflects where time or memory is going. It is statistical, not precise — like polling 1,000 voters and extrapolating to the entire population.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

pprof uses the OS signal mechanism (SIGPROF on Unix) for CPU profiling. When enabled, the runtime installs a signal handler that receives SIGPROF at the configured frequency (default 100Hz). On each signal, the handler calls sigprof which reads the interrupted goroutine's PC (program counter), SP (stack pointer), and FP (frame pointer) from the signal context. It unwinds the stack frame-by-frame using Go's stack unwinding metadata (gopclntab). The resulting call stack is hashed and stored in a hash table with an incrementing count. Every 10ms, one call stack is recorded. After 30 seconds, 3,000 call stacks form a statistically significant picture. For heap profiling, the sampling happens on allocation: each call to runtime.mallocgc randomly samples 1 allocation per 512KB based on MemProfileRate, recording the allocation stack trace and byte count.

## Run Instructions

```bash
go run ./curriculum/modules/13-performance-memory-engineering/lessons/01-pprof
go test ./curriculum/modules/13-performance-memory-engineering/lessons/01-pprof
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Profiling in production without CPU or memory overhead limits — pprof sampling adds ~5% CPU overhead by default; leaving profiling on 24/7 at 100Hz on every instance can cost noticeable capacity. Use runtime.SetCPUProfileRate(0) to disable, or enable profiling only during investigation windows.
- Confusing sampling with tracing — pprof samples every 100th event (default 100Hz CPU, 1 sample/512KB alloc). It does NOT capture every call or allocation. If a function runs 50 times between samples, only 1 appears in the profile. Use execution tracing (go test -trace, runtime/trace) when you need every event, not just statistical representation.
- Reading a single profile in isolation — a CPU profile taken during idle time shows different hot spots than one taken under load. Always compare profiles: 'before vs after' or 'idle vs loaded'. One profile proves nothing.

## In Production

Production Go services at every scale use pprof for performance investigation. The standard library exposes profiles via net/http/pprof on admin ports. Kubernetes liveness probes can trigger heap profile collection when memory exceeds a threshold. Datadog's Go profiler agent pulls pprof endpoints periodically. Google's internal Go profiling guide requires a profile before any performance optimization is accepted.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-13-02`.
