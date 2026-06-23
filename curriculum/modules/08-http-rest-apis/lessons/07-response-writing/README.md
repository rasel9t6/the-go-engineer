# Response writing

## Mission

Understand and apply Response writing in the context of professional Go software engineering.

## Prerequisites

- core-08-06

## Mental Model

Response writing follows a strict order: status code, headers, body. Go's ResponseWriter enforces this at the protocol level. Once you write the header or body, the status code is locked and cannot change.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Under the hood, http.ResponseWriter interacts with the Go runtime and operating system. Understanding this helps developers diagnose performance issues, debug unexpected errors, and write correct concurrent code. The key insight is that write HTTP responses with correct headers, status, and body is not magic — it follows well-defined Go language semantics and OS conventions.

## Run Instructions

```bash
go run ./curriculum/modules/08-http-rest-apis/lessons/07-response-writing
go test ./curriculum/modules/08-http-rest-apis/lessons/07-response-writing
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming write HTTP responses with correct headers, status, and body does something it does not — read the package docs instead of guessing behavior.
- Hardcoding Response writing values instead of parameterizing them — this makes tests impossible and changes painful.
- Copying Response writing code from StackOverflow without understanding the edge cases in your specific context.

## In Production

Professional Go engineers use Response writing daily in production services, CLI tools, API servers, and data pipelines. Understanding Response writing is essential for writing idiomatic, maintainable Go code that teams can confidently deploy and extend.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-08-08`.
