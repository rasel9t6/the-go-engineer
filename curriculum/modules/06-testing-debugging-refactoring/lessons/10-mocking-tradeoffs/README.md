# Mocking tradeoffs

## Mission

Understand and apply Mocking tradeoffs in the context of professional Go software engineering.

## Prerequisites

- core-06-09

## Mental Model

Mocking is like hiring a actor to play the role of a dependency in a play (test). The actor follows a script (expectations) and you verify they performed correctly (verify). But the actor only pretends -- there's no real implementation behind it.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Mocking frameworks generate types that implement the target interface. They track expected calls via an expectation chain. `EXPECT().Method(args).Return(values)` builds the expectation. `Ctrl.Finish()` verifies all expectations were met.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/10-mocking-tradeoffs
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/10-mocking-tradeoffs
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using mocks for every dependency -- including trivial ones.
- Creating brittle tests that break when implementation details change.
- Over-specifying mock expectations (exact call counts, exact parameters) when not needed.
- Using mocks instead of real implementations for stable, fast dependencies.

## In Production

Mocks are used when testing interaction patterns: verifying that a service calls a repository in the right order, that a handler sends the correct HTTP request, or that a retry mechanism correctly re-calls on failure.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-11`.
