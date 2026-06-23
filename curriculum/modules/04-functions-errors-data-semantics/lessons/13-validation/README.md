# Validation

## Mission

Understand and apply Validation in the context of professional Go software engineering.

## Prerequisites

- core-04-12

## Mental Model

Validation is a gatekeeper at the system boundary. Raw input arrives unchecked; the gatekeeper inspects it and either passes it through or rejects it with a reason. Garbage in, not passed through.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Struct tag validators use reflection to read struct tags and iterate over fields. For each field, they apply the registered validation function. Custom validators can access the full struct for cross-field validation.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/13-validation
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/13-validation
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Validating only at the UI layer and not in the domain/business logic layer.
- Returning `error` from validation but not providing which field failed or why.
- Using magic numbers or hardcoded strings for validation rules instead of constants or config.
- Performing validation after the operation has already started modifying state.

## In Production

Every web framework validates incoming JSON. Every database layer validates SQL queries. Every CLI validates flags and arguments. Validation is the first step in every non-trivial Go function that accepts external input.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-14`.
