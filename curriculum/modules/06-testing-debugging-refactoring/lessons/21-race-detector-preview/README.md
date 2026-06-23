# Race detector preview

## Mission

Understand and apply Race detector preview in the context of professional Go software engineering.

## Prerequisites

- core-06-20

## Mental Model

The race detector is like a security camera watching your program for simultaneous, unsynchronized access to the same memory. When two goroutines access the same variable without synchronization, the camera captures the moment and shows you the video (stack traces).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The race detector (ThreadSanitizer) instruments every memory access in the compiled program. It maintains a shadow memory tracking which goroutine last accessed each memory location. When two goroutines access the same location without synchronization, it reports the conflict.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/21-race-detector-preview
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/21-race-detector-preview
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Ignoring the race detector's warnings -- assuming they're false positives.
- Not testing with -race during development.
- Adding mutexes everywhere without understanding the actual race condition.
- Only running the race detector in CI and not during development.

## In Production

The race detector is used in development and CI for every concurrent Go program. It's particularly important for services, workers, and any program using goroutines with shared memory.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `module checkpoint`.
