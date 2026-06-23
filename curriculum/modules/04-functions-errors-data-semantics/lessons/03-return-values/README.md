# Return values

## Mission

Understand and apply Return values in the context of professional Go software engineering.

## Prerequisites

- core-04-02

## Mental Model

Return values are the function's output channels. The function signature declares what types of values it produces, and the caller receives those values.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go passes return values on the stack or in registers depending on the architecture and compiler version. For multiple return values, the compiler treats them as a hidden aggregate -- each return value occupies its own slot in the caller's frame. Named return values are just local variables that happen to be returned automatically on `return`.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/03-return-values
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/03-return-values
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Declaring return types but forgetting the `return` statement -- causes a compile error in Go.
- Using unnamed return values and still using `return` with no values -- leads to zero values being returned silently.
- Assuming a function with a return type but no `return` will panic at runtime -- it actually won't compile.

## In Production

Go's idiomatic `result, err` pattern is used in virtually every production Go codebase. Database `QueryRow` returns `(Row, error)`, HTTP clients return `(Response, error)`, and JSON functions return `(Output, error)`. This convention makes error handling explicit and chaining natural.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-04`.
