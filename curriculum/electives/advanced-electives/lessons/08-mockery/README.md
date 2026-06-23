# Mockery

## Mission

Understand and apply Mockery in the context of professional Go software engineering.

## Prerequisites

- elective-07

## Mental Model

Mockery reads a Go interface and generates a struct that implements it, with methods for setting expectations (On, Return), verifying calls (AssertExpectations), and tracking call count. The generated mock can be configured to return specific values, panic, or call through to a real implementation. Tests use the mock in place of the real dependency and assert that expected methods were called with the correct arguments.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Mockery uses go/packages to load and parse Go source files, extracts interface definitions, and generates a Go source file with a mock struct that implements the interface. Each mock method records call arguments and returns configured values. Expectations are stored in a map keyed by method name + argument values. The generated code uses sync.Mutex for thread-safe expectation tracking.

## Run Instructions

```bash
Read the lesson and complete the practice task.
No automated test is required for this lesson.
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Testing implementation details — asserting that a specific method was called when only the result matters, making tests brittle to refactoring.
- Not calling AssertExpectations in tests — expectations are never verified, and the test passes even if the mock was never called.
- Over-mocking — mocking interfaces that could use real implementations, creating fragile tests that fail when mock expectations are incorrect.
- Setting expectations that are too strict — using .Once() and exact argument matchers when the implementation could call the method multiple times or with slightly different arguments.
- Not regenerating mocks when interfaces change — the generated mock becomes stale and does not compile with the updated interface.

## In Production

Mockery is used in every Go project that follows interface-based testing: repository pattern tests (mock the database interface), HTTP handler tests (mock the service interface), event processing tests (mock the publisher interface). Production Go services use Mockery with go:generate directives so mocks are always in sync with interfaces. CI runs go generate ./... to regenerate mocks on interface changes.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-09`.
