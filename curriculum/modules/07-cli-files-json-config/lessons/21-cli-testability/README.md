# CLI testability

## Mission

Understand and apply CLI testability in the context of professional Go software engineering.

## Prerequisites

- core-07-20

## Mental Model

The pipeline may produce errors, partial results, or side effects. Each stage is independent and testable. The contract is: given valid input, produce valid output or an explicit error.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Under the hood, testing with os.Args / flag.FlagSet interacts with the Go runtime and operating system. Understanding this helps developers diagnose performance issues, debug unexpected errors, and write correct concurrent code. The key insight is that write testable command-line tools is not magic — it follows well-defined Go language semantics and OS conventions.

## Run Instructions

```bash
go run ./curriculum/modules/07-cli-files-json-config/lessons/21-cli-testability
go test ./curriculum/modules/07-cli-files-json-config/lessons/21-cli-testability
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming write testable command-line tools does something it does not — read the package docs instead of guessing behavior.
- Hardcoding CLI testability values instead of parameterizing them — this makes tests impossible and changes painful.
- Copying CLI testability code from StackOverflow without understanding the edge cases in your specific context.

## In Production

Professional Go engineers use CLI testability daily in production services, CLI tools, API servers, and data pipelines. Understanding CLI testability is essential for writing idiomatic, maintainable Go code that teams can confidently deploy and extend.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `module checkpoint`.
