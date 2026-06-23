# Table-driven tests

## Mission

Understand and apply Table-driven tests in the context of professional Go software engineering.

## Prerequisites

- core-06-02

## Mental Model

A table-driven test is like a recipe book with multiple recipes. Each recipe (table row) has ingredients (inputs), expected result, and a name. The cook (test function) follows each recipe and checks that the result matches expectations.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Table-driven tests work because Go's testing framework runs Test functions. Inside the function, you control the loop and reporting. `t.Run` creates a subtest that appears in the test output hierarchy.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/03-table-driven-tests
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/03-table-driven-tests
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Writing separate test functions for each test case instead of using table-driven tests.
- Not including a test name in each table entry -- failures are hard to identify.
- Using the same test inputs as the code -- tests should cover edge cases.
- Forgetting to add new cases to the table when adding new code paths.

## In Production

Table-driven tests are the idiomatic Go testing pattern. They are used in the standard library, open source projects, and production codebases. Code reviews expect table-driven tests for functions with multiple cases.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-04`.
