# Input validation for security

## Mission

Understand and apply Input validation for security in the context of professional Go software engineering.

## Prerequisites

- core-10-02

## Mental Model

Input validation is a series of gates: type check -> format check -> length check -> business rule check. Each gate passes data through or rejects it. Data is not trusted until it has passed every gate.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

encoding/json validates JSON structure (types, syntax). Additional validation uses struct tags or explicit Validate() methods. go-playground/validator uses reflection for struct tags.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/03-input-validation-for-security
go test ./curriculum/modules/10-auth-security/lessons/03-input-validation-for-security
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Validating input format but not semantics — checking email has @ but not that the domain exists.
- Validating input on the client side only — client validation is for UX, server validation is for security.
- Rejecting input that is 'too long' but accepting any length below the maximum without considering memory impact.

## In Production

Every production Go API validates input at the handler boundary. Frameworks like Echo and Gin provide built-in validation; net/http services use middleware or per-handler validation.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-04`.
