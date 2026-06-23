# net/http basics

## Mission

Understand and apply net/http basics in the context of professional Go software engineering.

## Prerequisites

- core-08-01

## Mental Model

net/http provides an HTTP client and server with a Handler interface at its core. Every server request enters through a handler, and every client call produces a response. The ServeMux routes requests to handlers using pattern matching on the request path.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Under the hood, http.ListenAndServe / http.HandleFunc interacts with the Go runtime and operating system. Understanding this helps developers diagnose performance issues, debug unexpected errors, and write correct concurrent code. The key insight is that start an HTTP server and handle requests is not magic — it follows well-defined Go language semantics and OS conventions.

## Run Instructions

```bash
go run ./curriculum/modules/08-http-rest-apis/lessons/02-net-http-basics
go test ./curriculum/modules/08-http-rest-apis/lessons/02-net-http-basics
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming start an HTTP server and handle requests does something it does not — read the package docs instead of guessing behavior.
- Hardcoding net/http basics values instead of parameterizing them — this makes tests impossible and changes painful.
- Copying net/http basics code from StackOverflow without understanding the edge cases in your specific context.

## In Production

Professional Go engineers use net/http basics daily in production services, CLI tools, API servers, and data pipelines. Understanding net/http basics is essential for writing idiomatic, maintainable Go code that teams can confidently deploy and extend.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-08-03`.
