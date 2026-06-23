# Debugging tests

## Mission

Understand and apply Debugging tests in the context of professional Go software engineering.

## Prerequisites

- core-06-17

## Mental Model

Debugging tests with Delve is like having a time machine for your test execution. You can pause at any test step, look at the state, and step forward to see how things evolve -- without modifying your test code.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

`dlv test` compiles the test binary with debug information, then executes it under Delve's control. The `--` separates Delve arguments from test arguments. `-test.run` is passed to the Go test runner to select a specific test.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/18-debugging-tests
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/18-debugging-tests
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using fmt.Println debugging when Delve would be faster.
- Not knowing that `dlv test` exists for debugging tests.
- Trying to debug a passing test instead of a failing one.
- Running the full test suite when only one test needs debugging.

## In Production

Developers routinely use `dlv test` to debug test failures. It's especially useful for table-driven tests where multiple cases run -- set breakpoints inside the loop to inspect specific cases.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-19`.
