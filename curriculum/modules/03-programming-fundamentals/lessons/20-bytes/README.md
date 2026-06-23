# Bytes

## Mission

Understand and apply Bytes in the context of professional Go software engineering.

## Prerequisites

- core-03-19

## Mental Model

A byte slice is mutable sequence of raw 8-bit values. Workbench for rearranging bits. Strings = framed photographs (look only). Byte slices = loose photos (crop, recolor freely).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

[]byte is regular slice with element type byte (uint8). bytes package optimized: bytes.Index uses assembly-level Rabin-Karp. string to []byte allocates new memory because strings are immutable.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/20-bytes
go test ./curriculum/modules/03-programming-fundamentals/lessons/20-bytes
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Confusing []byte and string - same underlying data, different mutability and operation sets.
- Assuming []byte auto-encodes UTF-8 - []byte is just bytes, no encoding implied.
- Using string concatenation when []byte append is more efficient.

## In Production

All Go I/O operates on []byte. Reading files, writing HTTP responses, encrypting, encoding JSON - all pass through byte slices.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-21`.
