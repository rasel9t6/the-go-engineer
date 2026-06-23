# Types

## Mission

Understand and apply Types in the context of professional Go software engineering.

## Prerequisites

- core-03-02

## Mental Model

A type is a blueprint specifying: (1) valid values, (2) permitted operations, (3) memory layout. int describes a number you can add; struct describes grouped fields; interface describes method sets.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Every Go type has a fixed compile-time size. Type descriptors store name, size, alignment, method set. Interface values carry concrete type pointer + data pointer. Type system is nominal, not structural.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/03-types
go test ./curriculum/modules/03-programming-fundamentals/lessons/03-types
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Confusing int and int64 - same size on 64-bit but compiler treats as distinct types.
- Assuming type alias allows mixing in operations without explicit conversion.
- Believing zero value of every type is nil - struct zero is all-zero fields, bool is false.

## In Production

Type safety is Go's primary guarantee. JSON unmarshaling, DB scanning, API parsing all rely on types. Understanding types means knowing what the compiler guarantees before code runs.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-04`.
