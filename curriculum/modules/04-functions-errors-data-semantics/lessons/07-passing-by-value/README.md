# Passing by value

## Mission

Understand and apply Passing by value in the context of professional Go software engineering.

## Prerequisites

- core-04-06

## Mental Model

Passing by value means the function gets its own copy of the argument -- like photocopying a document and handing the copy. For reference types (map, slice, pointer), the photocopy is of the address, so both copies point to the same underlying data.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The compiler determines whether to copy the argument to the stack or pass via registers. For `interface{}` parameters, the compiler generates an `eface` struct (type *uintptr, data unsafe.Pointer) and passes this two-word value.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/07-passing-by-value
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/07-passing-by-value
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Modifying a map inside a function and expecting the caller not to see the change -- maps are reference types.
- Passing a struct to a function, modifying it, and expecting the caller to see the modification.
- Thinking slices are 'passed by reference' -- the header is copied but the underlying array is shared.
- Using `&variable` as an argument when the parameter is a value, not a pointer.

## In Production

In production Go, pass by value is the default. When you read a config struct in an HTTP handler, you pass it by value to avoid accidental mutation. When you need to modify, you pass `*Config`. The caller always knows whether mutation is possible.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-08`.
