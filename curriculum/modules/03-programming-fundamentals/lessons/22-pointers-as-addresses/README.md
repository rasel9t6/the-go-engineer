# Pointers as addresses

## Mission

Understand and apply Pointers as addresses in the context of professional Go software engineering.

## Prerequisites

- core-03-21

## Mental Model

A pointer is a memory address - signpost saying 'data over here.' Instead of carrying data (copying), carry address (pointer). Multiple signposts can point to same house (aliasing).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go pointers are 8 bytes (64-bit) with virtual memory address. Pointer arithmetic forbidden. Dereferencing nil panics. new(T) returns *T. Method receivers distinguish value vs pointer - pointer allows mutation, avoids copy. Escape analysis determines stack vs heap.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/22-pointers-as-addresses
go test ./curriculum/modules/03-programming-fundamentals/lessons/22-pointers-as-addresses
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming pointer arithmetic works - Go does not support pointer arithmetic.
- Thinking pointers are only for large structs - pointers create aliases for shared access.
- Confusing *T (pointer to T) with *p dereference - context determines meaning.

## In Production

Method receivers use pointers. Large structs passed as pointers. DB rows, JSON decoders return pointers. Pointers essential for shared mutable state.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-23`.
