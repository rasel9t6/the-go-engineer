# Zero values

## Mission

Understand and apply Zero values in the context of professional Go software engineering.

## Prerequisites

- core-03-03

## Mental Model

Every type has a default starting state. Like a new house: int = room with 0 on wall, string = empty sign, pointer = no address. Zero value ensures no uninitialized gaps.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go compiler ensures every declared variable occupies memory. Without initializer, compiler emits zero-fill instruction. Pointers get nil. Key insight: zero values are not uninitialized - they are explicitly initialized to a known default.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/04-zero-values
go test ./curriculum/modules/03-programming-fundamentals/lessons/04-zero-values
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming pointers default to valid address instead of nil - dereferencing nil panics.
- Thinking struct zero values are nil - struct zero has all fields at zero, valid and usable.
- Forgetting map, chan, slice zero values are nil and cannot be written without init.

## In Production

Every var without initializer relies on zero values. JSON reading, handler init, data pipelines all depend on knowing what zero value guarantees.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-05`.
