# Debugging panics

## Mission

Understand and apply Debugging panics in the context of professional Go software engineering.

## Prerequisites

- core-06-12

## Mental Model

A panic is like a fire alarm in a building. Normal execution stops (unwinding). People evacuate (defers run). The fire marshal (recover) can declare it a false alarm, and the building returns to normal operation. Without recovery, the building burns down (program crashes).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The runtime tracks panicking goroutines. When a panic occurs, the runtime walks the defer chain. For each deferred function, if it calls recover(), the panic state is cleared and the goroutine resumes normal execution. Otherwise, the unwinding continues.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/13-debugging-panics
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/13-debugging-panics
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Not handling panic recovery at appropriate boundaries (goroutine boundaries, API boundaries).
- Recovering from a panic and continuing as if nothing happened -- the program state may be corrupted.
- Calling recover() outside a deferred function -- it returns nil.
- Not logging the panic details before recovering.

## In Production

HTTP servers recover from handler panics to return 500 instead of crashing. Long-running goroutines in background workers recover to continue processing. Database drivers recover from panics without crashing the application.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-14`.
