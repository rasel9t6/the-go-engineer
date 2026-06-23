# Composition

## Mission

Understand and apply Composition in the context of professional Go software engineering.

## Prerequisites

- core-05-05

## Mental Model

Composition is like LEGO bricks. Build small, single-purpose bricks (types) and snap them together. Each brick handles its responsibility. The composed object has all brick capabilities through method promotion.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

When A embeds B, the compiler inlines B's fields into A's memory layout. B's methods are promoted to A's method set via auto-generated forwarding methods. This is resolved at compile time.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/06-composition
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/06-composition
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Trying to use inheritance patterns in Go -- Go uses composition, not inheritance.
- Creating deep nesting of structs inside structs -- prefer flat composition.
- Using embedding when a named field would be more explicit.
- Forgetting that interfaces compose too -- different from struct composition.

## In Production

Production Go codebases compose service structs: `type UserService struct { db *sql.DB; logger *zap.Logger; cache *redis.Client }`. APIs compose middleware: `auth(rateLimit(handler))`.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-07`.
