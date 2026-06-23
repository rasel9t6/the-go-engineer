# Integration tests

## Mission

Understand and apply Integration tests in the context of professional Go software engineering.

## Prerequisites

- core-09-24

## Mental Model

The pipeline may produce errors, partial results, or side effects. Each stage is independent and testable. The contract is: given valid input, produce valid output or an explicit error.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Under the hood, testcontainers or Docker pattern interacts with the Go runtime and operating system. Understanding this helps developers diagnose performance issues, debug unexpected errors, and write correct concurrent code. The key insight is that write database integration tests is not magic — it follows well-defined Go language semantics and OS conventions.

## Run Instructions

```bash
go run ./curriculum/modules/09-sql-postgres-persistence/lessons/25-integration-tests
go test ./curriculum/modules/09-sql-postgres-persistence/lessons/25-integration-tests
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming write database integration tests does something it does not — read the package docs instead of guessing behavior.
- Hardcoding Integration tests values instead of parameterizing them — this makes tests impossible and changes painful.
- Copying Integration tests code from StackOverflow without understanding the edge cases in your specific context.

## In Production

Professional Go engineers use Integration tests daily in production services, CLI tools, API servers, and data pipelines. Understanding Integration tests is essential for writing idiomatic, maintainable Go code that teams can confidently deploy and extend.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-09-26`.
