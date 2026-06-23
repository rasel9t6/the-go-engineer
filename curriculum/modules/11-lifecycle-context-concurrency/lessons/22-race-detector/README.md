# Race detector

## Mission

Understand and apply Race detector in the context of professional Go software engineering.

## Prerequisites

- core-11-21

## Mental Model

It establishes a contract between goroutines about who does what, when, and in what order. Correct usage follows the Go proverb: 'Don't communicate by sharing memory; share memory by communicating.'

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Under the hood, race detector in Go is implemented using the runtime's synchronization primitives (mutex, semaphore, atomic operations). The runtime scheduler coordinates goroutine blocking and waking to minimize overhead.

## Run Instructions

```bash
go run ./curriculum/modules/11-lifecycle-context-concurrency/lessons/22-race-detector
go test ./curriculum/modules/11-lifecycle-context-concurrency/lessons/22-race-detector
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Applying race detector without understanding the concurrency model — the wrong primitive causes deadlocks or data races.
- Assuming race detector works the same as similar primitives in other languages — Go's channel and mutex semantics differ from Java/C++.
- Copying race detector patterns from tutorials without adapting them to the specific synchronization needs of the application.

## In Production

Production Go services use race detector in request handling, background processing, metrics collection, and graceful shutdown. Understanding this primitive is essential for writing production-grade concurrent Go code.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-11-23`.
