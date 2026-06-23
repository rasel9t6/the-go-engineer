# Test fixtures

## Mission

Understand and apply Test fixtures in the context of professional Go software engineering.

## Prerequisites

- core-06-05

## Mental Model

Test fixtures are like photographs of expected test inputs and outputs. You keep the photo album (testdata directory) with your tests. Each test opens the relevant photo and checks current output against it.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The `go test` command changes the working directory to the package directory before running tests. This means relative paths to testdata resolve correctly regardless of where `go test` is invoked from. The testdata directory itself is skipped by the Go compiler.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/06-test-fixtures
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/06-test-fixtures
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Creating test fixtures inside test functions instead of using testdata/ directory.
- Modifying test fixtures during tests -- they should be read-only.
- Not using t.TempDir() for temporary test directories.
- Hardcoding file paths that break when tests are run from different working directories.

## In Production

Production Go tests use fixtures for JSON input samples, SQL migration files, configuration files, and binary test data. The testdata directory is a Go convention recognized by the toolchain.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-07`.
