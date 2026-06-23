# Value receivers

## Mission

Understand and apply Value receivers in the context of professional Go software engineering.

## Prerequisites

- core-05-03

## Mental Model

A value receiver is like taking a photograph and describing it. You can talk about it, write on the photo, but the original object stays unchanged. The photo is a snapshot at call time.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

When a value receiver method is called on a pointer, the compiler inserts a `*ptr` dereference and copies the value. The copy is on the caller's stack if small enough, or on the heap if escape analysis says so.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/04-value-receivers
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/04-value-receivers
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using a value receiver for a large struct and wondering why the method is slow.
- Assuming a value receiver method can modify the original -- it receives a copy.
- Using value receivers for methods that should handle nil receivers.
- Forgetting value receiver methods work on both pointer and value types.

## In Production

`time.Time` uses value receivers exclusively. Small data types like `net.IP` and value objects use value receivers. Use value receivers for immutability semantics.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-05`.
