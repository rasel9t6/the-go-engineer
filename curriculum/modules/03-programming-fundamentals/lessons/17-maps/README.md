# Maps

## Mission

Understand and apply Maps in the context of professional Go software engineering.

## Prerequisites

- core-03-16

## Mental Model

A map is a hash table - key-value pairs where keys hash to storage location. Like coat check: give key, get ticket (hash), coat stored in numbered booth.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A map is pointer to runtime's hmap struct. make allocates hmap and bucket array. Key hashed using per-map random seed, placed in bucket. Missing key returns zero value.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/17-maps
go test ./curriculum/modules/03-programming-fundamentals/lessons/17-maps
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Reading nil map returns zero - no panic. Writing nil map panics.
- Assuming map iteration order is deterministic - Go randomizes intentionally.
- Using slices as map keys - slices not comparable. Use strings, ints, or struct keys.

## In Production

Every API handler parses query params into maps. Every session store is a map. Maps are primary associative data structure.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-18`.
