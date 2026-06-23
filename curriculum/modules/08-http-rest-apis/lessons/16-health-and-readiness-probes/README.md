# Health and readiness probes

## Mission

Understand and apply Health and readiness probes in the context of professional Go software engineering.

## Prerequisites

- core-08-15

## Mental Model

Health probes report whether the process is alive. Readiness probes report whether the server can accept traffic — databases warmed, caches loaded, dependencies reachable. Separate endpoints let orchestration systems make informed decisions.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Under the hood, health endpoint pattern interacts with the Go runtime and operating system. Understanding this helps developers diagnose performance issues, debug unexpected errors, and write correct concurrent code. The key insight is that implement health check endpoints for orchestration is not magic — it follows well-defined Go language semantics and OS conventions.

## Run Instructions

```bash
go run ./curriculum/modules/08-http-rest-apis/lessons/16-health-and-readiness-probes
go test ./curriculum/modules/08-http-rest-apis/lessons/16-health-and-readiness-probes
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming implement health check endpoints for orchestration does something it does not — read the package docs instead of guessing behavior.
- Hardcoding Health and readiness probes values instead of parameterizing them — this makes tests impossible and changes painful.
- Copying Health and readiness probes code from StackOverflow without understanding the edge cases in your specific context.

## In Production

Professional Go engineers use Health and readiness probes daily in production services, CLI tools, API servers, and data pipelines. Understanding Health and readiness probes is essential for writing idiomatic, maintainable Go code that teams can confidently deploy and extend.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-08-17`.
