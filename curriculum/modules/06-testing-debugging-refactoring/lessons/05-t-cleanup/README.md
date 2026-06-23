# t.Cleanup

## Mission

Understand and apply t.Cleanup in the context of professional Go software engineering.

## Prerequisites

- core-06-04

## Mental Model

t.Cleanup is like checking luggage at an airport. You hand over items (register cleanup) as you acquire them. When you leave (test completes), items are returned in reverse order (LIFO). This ensures proper resource release order.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

t.Cleanup adds a cleanup function to an internal stack in the test's internal state. When the test completes (pass, fail, or panic), the stack is popped in LIFO order. Cleanups registered in subtests run when the subtest completes.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/05-t-cleanup
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/05-t-cleanup
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Forgetting that t.Cleanup runs in LIFO order, not FIFO.
- Using defer for cleanup instead of t.Cleanup -- defer runs on function return, t.Cleanup runs when the test and all its subtests complete.
- Registering cleanup in a helper function that's called from multiple tests -- cleanup registers per test.
- Not understanding that t.Cleanup can be called multiple times.

## In Production

t.Cleanup is used for cleaning up temp directories, closing test databases, stopping test servers, and restoring global state. It's the standard teardown mechanism in modern Go tests.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-06`.
