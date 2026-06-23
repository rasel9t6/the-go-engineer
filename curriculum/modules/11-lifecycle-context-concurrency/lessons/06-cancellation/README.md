# Cancellation

## Mission

Understand and apply Cancellation in the context of professional Go software engineering.

## Prerequisites

- core-11-05

## Mental Model

Context cancellation is a tree broadcast. The root context branches into child contexts — when any parent is canceled, all children receive the cancellation. The done channel is like a fire alarm: once it rings, it keeps ringing, and every goroutine in the building should have a plan for what to do when it hears it.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

context.WithCancel creates a child context and a cancel function. Internally, a new cancelCtx struct is created with a channel (ctx.Done()) — initially nil. When cancel() is called, the channel is closed (or created and immediately closed), and the canceler iterates over its list of child contexts, calling cancel() on each. context.WithTimeout wraps WithCancel with a time.AfterFunc that calls cancel() after the deadline — the timer is stored so that cancel() can stop the timer if called before the deadline. The ctx.Err() method returns nil before cancellation, Canceled after cancel() is called, or DeadlineExceeded after the timeout fires.

## Run Instructions

```bash
go run ./curriculum/modules/11-lifecycle-context-concurrency/lessons/06-cancellation
go test ./curriculum/modules/11-lifecycle-context-concurrency/lessons/06-cancellation
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Calling context.WithTimeout in every function instead of accepting a context from the caller — creates fragmented cancellation trees that are impossible to trace.
- Ignoring the return value of context.WithTimeout and using the parent context in the goroutine — the goroutine never receives cancellation.
- Checking ctx.Err() after the fact but never selecting on ctx.Done() in a goroutine — the goroutine blocks forever because it never listens for cancellation.
- Passing a context to a function that stores it in a struct — the context outlives the request, causing stale cancellation state to affect unrelated operations.

## In Production

Production Go services use context cancellation in every HTTP handler, every database query, every RPC call, and every goroutine spawn. Without correct cancellation, a simple deployment rollback can leave hundreds of orphaned goroutines consuming memory and holding database connections.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-11-07`.
