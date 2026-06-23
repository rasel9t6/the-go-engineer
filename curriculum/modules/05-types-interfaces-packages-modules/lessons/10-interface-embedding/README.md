# Interface embedding

## Mission

Understand and apply Interface embedding in the context of professional Go software engineering.

## Prerequisites

- core-05-09

## Mental Model

Interface embedding is like a multi-tool combining simple tools. A `ReadWriter` is a Swiss Army knife with knife (Reader) and scissors (Writer). Use as either, and satisfies any interface needing either.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The compiler builds the method table by concatenating all embedded interface methods. Duplicate methods must have identical signatures. The itab entry points to the concrete implementation.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/10-interface-embedding
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/10-interface-embedding
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Embedding an interface in a struct and forgetting to provide an implementation -- nil pointer panic.
- Creating circular dependencies when embedding interfaces across packages.
- Using interface embedding when explicit named fields would be clearer.
- Embedding too many interfaces into one monolithic combined interface.

## In Production

Production Go uses `io.ReadWriter` for buffered I/O. Database drivers compose Querier, Executor, Transactor. `http.ResponseWriter` embeds `io.Writer`. Interface embedding builds complex contracts from simple parts.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-11`.
