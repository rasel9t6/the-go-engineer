# Range

## Mission

Understand and apply Range in the context of professional Go software engineering.

## Prerequisites

- core-03-11

## Mental Model

Range is a universal iterator that adapts to collection type. Like universal remote - pressing 'next' gives index+value on array, key+value on map, byte-index+rune on string.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Range is a compiler primitive. For slices: simple index loop. For maps: iterator over hash table buckets, starting random. For strings: rune-decoding iterator. Range variable is always a copy.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/12-range
go test ./curriculum/modules/03-programming-fundamentals/lessons/12-range
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Forgetting second range return - for i := range gives only index, use for i, v := range.
- Assuming range over map yields insertion order - Go randomizes map iteration.
- Modifying range value (v) thinking it updates slice - v is a copy.

## In Production

Processing HTTP headers, scanning file contents, decoding request bodies, iterating query results all use range.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-13`.
