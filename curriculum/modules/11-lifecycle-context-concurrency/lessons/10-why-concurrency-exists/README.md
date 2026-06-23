# Why concurrency exists

## Mission

Understand and apply Why concurrency exists in the context of professional Go software engineering.

## Prerequisites

- core-11-09

## Mental Model

It establishes a contract between goroutines about who does what, when, and in what order. Correct usage follows the Go proverb: 'Don't communicate by sharing memory; share memory by communicating.'

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Under the hood, why concurrency exists in Go is implemented using the runtime's synchronization primitives (mutex, semaphore, atomic operations). The runtime scheduler coordinates goroutine blocking and waking to minimize overhead.

## Run Instructions

```bash
go run ./curriculum/modules/11-lifecycle-context-concurrency/lessons/10-why-concurrency-exists
go test ./curriculum/modules/11-lifecycle-context-concurrency/lessons/10-why-concurrency-exists
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Applying why concurrency exists without understanding the concurrency model — the wrong primitive causes deadlocks or data races.
- Assuming why concurrency exists works the same as similar primitives in other languages — Go's channel and mutex semantics differ from Java/C++.
- Copying why concurrency exists patterns from tutorials without adapting them to the specific synchronization needs of the application.

## In Production

Production Go services use why concurrency exists in request handling, background processing, metrics collection, and graceful shutdown. Understanding this primitive is essential for writing production-grade concurrent Go code.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-11-11`.
