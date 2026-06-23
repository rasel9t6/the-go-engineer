# Multiple return values

## Mission

Understand and apply Multiple return values in the context of professional Go software engineering.

## Prerequisites

- core-04-03

## Mental Model

A function with multiple return values is like a machine with multiple output chutes. Each chute produces a specific type of output. The caller must be ready to catch all outputs -- Go does not allow ignoring return values.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The compiler treats multiple returns as an implicit struct -- each return value occupies a slot in the caller's stack frame. The return statement packs all values into their respective slots atomically. On amd64, small returns use registers; larger returns spill to the stack.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/04-multiple-return-values
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/04-multiple-return-values
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Swapping the order of return values -- convention is result first, error last.
- Using multiple return values for unrelated outputs instead of returning a struct.
- Calling a function but only capturing the first return value with `_` for the error.
- Forgetting to handle the error because multiple returns make it easy to ignore.

## In Production

Go's `result, err` pattern is the standard for all I/O operations. `db.QueryRow(...).Scan(&id, &name)` uses multiple returns from Scan. `json.Marshal` returns `([]byte, error)`. The convention makes it impossible to accidentally use a result when the operation failed.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-05`.
