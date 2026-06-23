# Goroutines

## Mission

Understand and apply Goroutines in the context of professional Go software engineering.

## Prerequisites

- core-11-10

## Mental Model

A goroutine is a lightweight thread managed by the Go runtime. Unlike OS threads (~1MB stack, expensive context switch), a goroutine starts with a tiny stack (~4KB, now ~2KB minimum) and is multiplexed onto OS threads by the Go scheduler. The go keyword forks the current execution path: the caller continues immediately, and the new goroutine starts executing the function on its own stack. Goroutines are cheap enough to create by the thousands but not free — each consumes memory and scheduling resources.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The go keyword compiles to a call to runtime.newproc. runtime.newproc allocates a goroutine struct (g) containing the stack (initially ~2KB), the instruction pointer for the function, and signal/defer stacks. The g is added to the current P's run queue. The Go scheduler runs as part of the runtime's scheduler loop (schedule function), which picks runnable goroutines from the P's local run queue or steals from other Ps. Each M (OS thread) runs a P's goroutines in a tight loop: execute, block, schedule next. When a goroutine blocks on I/O or a syscall, the M is released to run other goroutines, and a new M may be created to handle the blocking call. The scheduler is cooperative-preemptive: a goroutine runs until it makes a function call (where the scheduler can check for preemption) or blocks — it does not use OS-level preemption (except for signals used for async preemption since Go 1.14).

## Run Instructions

```bash
go run ./curriculum/modules/11-lifecycle-context-concurrency/lessons/11-goroutines
go test ./curriculum/modules/11-lifecycle-context-concurrency/lessons/11-goroutines
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Starting a goroutine and immediately expecting its result — goroutines execute concurrently, not sequentially; the main function may exit before the goroutine finishes, silently discarding the result.
- Assuming goroutines have zero startup cost — creating 100,000 goroutines instantly schedules 100,000 goroutines on the run queue, each with its own stack, causing a burst of scheduling overhead and memory allocation.
- Using time.Sleep to wait for a goroutine to finish — Sleep is not synchronization; it works coincidentally under light load and fails under heavy load when the goroutine has not had time to run.
- Sharing mutable state between goroutines without synchronization — two goroutines writing to the same map or slice without a mutex causes a data race that corrupts memory and produces non-deterministic crashes.

## In Production

Every Go HTTP server creates a goroutine per connection — a server with 10,000 concurrent connections has 10,000 goroutines. Background workers use goroutines for periodic tasks (log rotation, cache refresh). Pipeline processing uses goroutines connected by channels for producer-consumer parallelism. Worker pools (using an unbuffered channel) limit goroutine count to GOMAXPROCS for CPU-bound work.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-11-12`.
