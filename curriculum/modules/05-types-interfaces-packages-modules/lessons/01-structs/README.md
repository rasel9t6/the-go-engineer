# Structs

## Mission

Understand and apply Structs in the context of professional Go software engineering.

## Prerequisites

- core-04-18

## Mental Model

A struct is a container that groups related data fields into a single value. The struct bundles these together as one unit.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A struct is laid out as a contiguous block of its fields in declaration order. The compiler adds alignment padding between fields. Reordering fields from largest to smallest alignment reduces padding.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/01-structs
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/01-structs
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Forgetting that struct fields are comma-separated -- missing a comma on multi-line declarations causes compile error.
- Declaring a struct type inside a function and trying to use it outside -- struct types are block-scoped.
- Comparing structs with `==` when they contain slices or maps -- incomparable types cause compile error.
- Assuming zero values of struct fields are nil for all types -- int fields are 0, bool fields are false.

## In Production

Structs are used in every Go codebase for domain models (`Order`, `User`, `Product`), configuration (`Config`), API request/response bodies (`CreateUserRequest`), and JSON messages.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-02`.
