# Middleware pattern

## Mission

Understand and apply Middleware pattern in the context of professional Go software engineering.

## Prerequisites

- core-08-09

## Mental Model

Middleware wraps a handler to add cross-cutting behaviour — logging, auth, rate limiting — without modifying the handler itself. Middleware calls the next handler in the chain, creating a before-and-after pipeline around the core logic.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Under the hood, http.Handler middleware pattern interacts with the Go runtime and operating system. Understanding this helps developers diagnose performance issues, debug unexpected errors, and write correct concurrent code. The key insight is that chain HTTP handlers for cross-cutting concerns is not magic — it follows well-defined Go language semantics and OS conventions.

## Run Instructions

```bash
go run ./curriculum/modules/08-http-rest-apis/lessons/10-middleware-pattern
go test ./curriculum/modules/08-http-rest-apis/lessons/10-middleware-pattern
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming chain HTTP handlers for cross-cutting concerns does something it does not — read the package docs instead of guessing behavior.
- Hardcoding Middleware pattern values instead of parameterizing them — this makes tests impossible and changes painful.
- Copying Middleware pattern code from StackOverflow without understanding the edge cases in your specific context.

## In Production

Professional Go engineers use Middleware pattern daily in production services, CLI tools, API servers, and data pipelines. Understanding Middleware pattern is essential for writing idiomatic, maintainable Go code that teams can confidently deploy and extend.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-08-11`.
