# Custom error types

## Mission

Understand and apply Custom error types in the context of professional Go software engineering.

## Prerequisites

- core-04-17

## Mental Model

A custom error type is like a specialized alert form. A sentinel error is a simple post-it note. A custom error type is a structured form with labeled fields carrying actionable information.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A custom error type satisfies `error` by implementing `Error() string`. The interface itself is a two-word struct (type pointer + data pointer). errors.As uses type assertion to extract the custom type.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/18-custom-error-types
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/18-custom-error-types
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Defining a custom error type when a simple sentinel error would suffice.
- Making the custom error type too complex -- errors should carry only what callers need.
- Forgetting to implement the `Error() string` method -- the type won't satisfy `error`.
- Comparing custom errors with `==` instead of `errors.Is` -- `==` fails when wrapped.

## In Production

API services define `APIError { Code int, Message string }` for structured JSON responses. Database layers define `DBError { Code string, Message string, Constraint string }`. Custom error types are standard for structured error data.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `module checkpoint`.
