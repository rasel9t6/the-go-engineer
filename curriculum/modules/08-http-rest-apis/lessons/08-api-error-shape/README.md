# API error shape

## Mission

Understand and apply API error shape in the context of professional Go software engineering.

## Prerequisites

- core-08-07

## Mental Model

API errors must carry a code, a message, and optionally details in a consistent structure. A uniform error shape lets clients parse failures programmatically instead of scraping response text.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Under the hood, error response struct pattern interacts with the Go runtime and operating system. Understanding this helps developers diagnose performance issues, debug unexpected errors, and write correct concurrent code. The key insight is that design consistent error response structures is not magic — it follows well-defined Go language semantics and OS conventions.

## Run Instructions

```bash
go run ./curriculum/modules/08-http-rest-apis/lessons/08-api-error-shape
go test ./curriculum/modules/08-http-rest-apis/lessons/08-api-error-shape
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming design consistent error response structures does something it does not — read the package docs instead of guessing behavior.
- Hardcoding API error shape values instead of parameterizing them — this makes tests impossible and changes painful.
- Copying API error shape code from StackOverflow without understanding the edge cases in your specific context.

## In Production

Professional Go engineers use API error shape daily in production services, CLI tools, API servers, and data pipelines. Understanding API error shape is essential for writing idiomatic, maintainable Go code that teams can confidently deploy and extend.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-08-09`.
