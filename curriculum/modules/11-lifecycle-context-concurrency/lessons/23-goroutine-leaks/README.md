# Goroutine leaks

## Mission

Understand and apply Goroutine leaks in the context of professional Go software engineering.

## Prerequisites

- core-11-22

## Mental Model

A goroutine is an independently scheduled execution unit that MUST have a guaranteed way to stop. If a goroutine cannot reach a return statement, it lives forever — even if blocked on a channel, a syscall, or a mutex. A goroutine leak is a permanent allocation: the goroutine's stack is never freed, and if it holds references to other objects, those objects are never GC'd either. The only way to avoid leaks is to ensure every goroutine has a clear exit path: context cancellation, channel close, or a done signal.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Goroutines are user-space threads scheduled by the Go runtime's M:P:G model. A leaked goroutine transitions from _Grunning to _Gwaiting (blocked on a channel, mutex, or select) or _Gsyscall (blocked on a system call). It stays in this state indefinitely. The runtime's allgs slice holds a pointer to every goroutine ever created — leaked goroutines are never removed. GC must scan each leaked goroutine's stack during every GC cycle, adding O(leaked) overhead. The pprof goroutine profile dumps all goroutine stacks with their state and blocking location — this is the primary tool for diagnosing leaks.

## Run Instructions

```bash
go run ./curriculum/modules/11-lifecycle-context-concurrency/lessons/23-goroutine-leaks
go test ./curriculum/modules/11-lifecycle-context-concurrency/lessons/23-goroutine-leaks
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Starting a goroutine without a shutdown signal — the goroutine runs forever because there is no channel close, context cancellation, or done channel to tell it to stop.
- Using time.Sleep as a goroutine cleanup mechanism — sleeping for 1 second before returning is not a shutdown strategy; it leaks for the sleep duration and leaves a window where the goroutine runs after the caller has returned.
- Sending on a channel when the receiver has already stopped — the send blocks forever because there is no reader, and the goroutine is permanently blocked.
- Assuming a goroutine stops when its parent function returns — the goroutine has its own stack and continues running independently until it exits or blocks permanently.

## In Production

Goroutine leaks are the #1 cause of memory growth in Go services (after heap allocation). A common production pattern: an HTTP handler starts a goroutine to send a metric or log to a background system — if the handler returns and the background system is slow, the goroutine blocks on a channel send forever. These accumulate under load and eventually OOM the server. Every Go team at scale (Uber, Lyft, Docker) has been hit by goroutine leaks in production.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-11-24`.
