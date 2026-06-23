# Constants

## Mission

Understand and apply Constants in the context of professional Go software engineering.

## Prerequisites

- core-03-05

## Mental Model

A constant is a compile-time known value that never changes. Like a sticky note with a number - use wherever needed, cannot modify. Untyped = flexible sticky note that adapts to any type.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go compiler evaluates constant expressions with a special arbitrary-precision engine. Untyped integers can store millions of digits. Default type inferred from representation: 42 = int, 3.14 = float64, 'x' = rune.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/06-constants
go test ./curriculum/modules/03-programming-fundamentals/lessons/06-constants
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming constants are variables that cannot change - they are compile-time values with no address.
- Thinking const x float64 = 3.14 creates runtime float64 - untyped constants have arbitrary precision at compile time.
- Confusing const grouping with iota - they work together but serve different purposes.

## In Production

Constants define HTTP status codes, time durations, config defaults, enum-like values. Go services rely on constants for error sentinels (io.EOF), permission modes, time arithmetic.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-07`.
