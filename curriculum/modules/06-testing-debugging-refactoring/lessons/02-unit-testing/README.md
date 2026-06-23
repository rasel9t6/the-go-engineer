# Unit testing

## Mission

Understand and apply Unit testing in the context of professional Go software engineering.

## Prerequisites

- core-06-01

## Mental Model

A unit test is like testing a single gear from a clock. You take the gear out, spin it with known force, and check that it rotates the expected amount. You don't care about the rest of the clock for this test.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's test runner discovers tests by looking for `func TestXxx(t *testing.T)` in _test.go files. Each test is run sequentially unless subtests are used. The -v flag shows verbose output including each test's result.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/02-unit-testing
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/02-unit-testing
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Writing integration tests instead of unit tests -- testing too many layers at once.
- Testing private functions through the public API -- acceptable but sometimes calls for internal tests.
- Not using table-driven tests for multiple test cases -- repetitive test code.
- Ignoring test output and not reading test failures.

## In Production

Unit tests are the foundation of Go testing. Every standard library package has unit tests. Most functions in Go production codebases have associated unit tests.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-03`.
