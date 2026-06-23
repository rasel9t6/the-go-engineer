# REST design principles

## Mission

Understand and apply REST design principles in the context of professional Go software engineering.

## Prerequisites

- core-08-16

## Mental Model

REST organises an API around resources identified by URLs, with actions expressed through HTTP methods: GET reads, POST creates, PUT replaces, PATCH updates, DELETE removes.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Under the hood, REST conventions in Go interacts with the Go runtime and operating system. Understanding this helps developers diagnose performance issues, debug unexpected errors, and write correct concurrent code. The key insight is that design RESTful APIs following resource-oriented design is not magic — it follows well-defined Go language semantics and OS conventions.

## Run Instructions

```bash
go run ./curriculum/modules/08-http-rest-apis/lessons/17-rest-design-principles
go test ./curriculum/modules/08-http-rest-apis/lessons/17-rest-design-principles
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming design RESTful APIs following resource-oriented design does something it does not — read the package docs instead of guessing behavior.
- Hardcoding REST design principles values instead of parameterizing them — this makes tests impossible and changes painful.
- Copying REST design principles code from StackOverflow without understanding the edge cases in your specific context.

## In Production

Professional Go engineers use REST design principles daily in production services, CLI tools, API servers, and data pipelines. Understanding REST design principles is essential for writing idiomatic, maintainable Go code that teams can confidently deploy and extend.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-08-18`.
