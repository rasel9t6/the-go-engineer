# Fakes before mocks

## Mission

Understand and apply Fakes before mocks in the context of professional Go software engineering.

## Prerequisites

- core-06-08

## Mental Model

A fake is like a flight simulator for testing pilots. It doesn't mock individual instrument readings -- it runs a simplified but complete version of the flight software. The pilot (test) can fly patterns without taking off.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A fake is a struct that implements an interface with simplified but functional logic. It doesn't use expectation libraries -- it just records operations and provides accessors for test assertions. Fakes are idiomatic Go because they rely on interfaces, which Go satisfies implicitly.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/09-fakes-before-mocks
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/09-fakes-before-mocks
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Reaching for mocking frameworks before trying Go interfaces with simple fakes.
- Creating overly complex mock setups that need maintenance.
- Using mocks for types that could be real implementations.
- Not understanding that fakes (simple implementations) are often better than mocks.

## In Production

Production Go codebases use fakes for databases (in-memory maps), HTTP clients (returning canned responses), and email senders (capturing sent messages). Fakes are the preferred test double approach.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-10`.
