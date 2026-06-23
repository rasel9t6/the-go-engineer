# Why testing exists

## Mission

Understand and apply Why testing exists in the context of professional Go software engineering.

## Prerequisites

- core-05-23

## Mental Model

Tests are a safety net. They catch regressions, document expected behavior, and make refactoring safe. The test suite is the executable specification of what the code should do. If the tests pass, the code works as intended.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's test runner compiles test files (those ending in _test.go) alongside the package. Each Test function is called with a *testing.T parameter. The test passes if no failure methods (t.Error, t.Fatal) are called before the function returns.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/01-why-testing-exists
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/01-why-testing-exists
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Treating tests as an afterthought -- writing them after the code is 'done'.
- Not understanding that tests are executable documentation.
- Believing manual testing is sufficient for correctness.
- Testing only happy paths and skipping error cases.

## In Production

Every production Go service has an extensive test suite. CI pipelines run tests on every pull request. Code review requires new tests for new code. Testing is not optional in professional Go development.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-02`.
