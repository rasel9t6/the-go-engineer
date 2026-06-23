# Request parsing

## Mission

Understand and apply Request parsing in the context of professional Go software engineering.

## Prerequisites

- core-08-04

## Mental Model

Parsing an HTTP request means reading the method, URL, headers, and body from the wire. Go provides Request.Form, Request.URL.Query, and json.Decoder for structured access. Each parsing strategy trades safety, performance, and ergonomics.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Under the hood, r.URL.Query / json.Decoder / r.FormValue interacts with the Go runtime and operating system. Understanding this helps developers diagnose performance issues, debug unexpected errors, and write correct concurrent code. The key insight is that parse HTTP request data into Go values is not magic — it follows well-defined Go language semantics and OS conventions.

## Run Instructions

```bash
go run ./curriculum/modules/08-http-rest-apis/lessons/05-request-parsing
go test ./curriculum/modules/08-http-rest-apis/lessons/05-request-parsing
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming parse HTTP request data into Go values does something it does not — read the package docs instead of guessing behavior.
- Hardcoding Request parsing values instead of parameterizing them — this makes tests impossible and changes painful.
- Copying Request parsing code from StackOverflow without understanding the edge cases in your specific context.

## In Production

Professional Go engineers use Request parsing daily in production services, CLI tools, API servers, and data pipelines. Understanding Request parsing is essential for writing idiomatic, maintainable Go code that teams can confidently deploy and extend.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-08-06`.
