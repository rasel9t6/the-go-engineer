# errors.As

## Mission

Understand and apply errors.As in the context of professional Go software engineering.

## Prerequisites

- core-04-11

## Mental Model

errors.As is like a detective looking for a specific person in a chain of nested rooms. Each room (error) might contain the target type. When found, the detective hands you a reference to that person.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

errors.As uses reflection-like assignment via the reflectlite package. It walks the error chain by calling `Unwrap()`. For each error, it checks if the error implements the target interface or if a concrete type can be assigned.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/12-errors-as
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/12-errors-as
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using `errors.As` when `errors.Is` would suffice -- As is for type extraction, Is is for value equality.
- Forgetting that `errors.As` requires a pointer to the target type.
- Using `errors.As(err, &target)` where target is already nil -- the assignment happens on match.
- Checking `errors.As` return value but then using target without verifying the assertion succeeded.

## In Production

Production Go uses errors.As to extract `*net.DNSError` for DNS server IP, `*pgconn.PgError` for PostgreSQL error codes, and custom domain errors with correlation IDs.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-13`.
