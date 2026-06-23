# Text templates

## Mission

Understand and apply Text templates in the context of professional Go software engineering.

## Prerequisites

- core-07-18

## Mental Model

The pipeline may produce errors, partial results, or side effects. Each stage is independent and testable. The contract is: given valid input, produce valid output or an explicit error.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Under the hood, text/template package interacts with the Go runtime and operating system. Understanding this helps developers diagnose performance issues, debug unexpected errors, and write correct concurrent code. The key insight is that render text output from templates is not magic — it follows well-defined Go language semantics and OS conventions.

## Run Instructions

```bash
go run ./curriculum/modules/07-cli-files-json-config/lessons/19-text-templates
go test ./curriculum/modules/07-cli-files-json-config/lessons/19-text-templates
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming render text output from templates does something it does not — read the package docs instead of guessing behavior.
- Hardcoding Text templates values instead of parameterizing them — this makes tests impossible and changes painful.
- Copying Text templates code from StackOverflow without understanding the edge cases in your specific context.

## In Production

Professional Go engineers use Text templates daily in production services, CLI tools, API servers, and data pipelines. Understanding Text templates is essential for writing idiomatic, maintainable Go code that teams can confidently deploy and extend.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-07-20`.
