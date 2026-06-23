# Embedding

## Mission

Understand and apply Embedding in the context of professional Go software engineering.

## Prerequisites

- core-05-06

## Mental Model

Embedding is like clicking two LEGO bricks together. The combined piece has all the studs (methods) of both bricks. Interior surfaces (fields) are accessible from outside. It's structural composition with automatic forwarding.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The compiler builds a flattened field layout. For `type A struct { B; x int }`, B's fields are at offset 0, then x. Method promotion generates forwarding methods in the outer type's method table.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/07-embedding
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/07-embedding
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Confusing embedding (struct-in-struct without field name) with composition via named fields.
- Expecting embedding to work like inheritance -- embedding promotes methods, not polymorphic behavior.
- Embedding the same type twice -- compile error.
- Forgetting that embedding a type exposes its methods through the outer type accidentally.

## In Production

Production services embed `sync.Mutex` for thread safety. API handlers embed `base.Handler` for common methods. `embed.FS` embeds file systems. Embedding is universal for compositional reuse.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-08`.
