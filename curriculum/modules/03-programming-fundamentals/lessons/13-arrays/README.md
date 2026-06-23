# Arrays

## Mission

Understand and apply Arrays in the context of professional Go software engineering.

## Prerequisites

- core-03-12

## Mental Model

An array is fixed-length contiguous memory. Like parking lot with numbered spaces - size determined at construction, cannot add/remove spaces. Copying copies the entire lot.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

var arr [5]int creates 5 * 8 = 40 bytes contiguous. Length baked into type. Value type: arr1 = arr2 copies all elements. Indices bounds-checked at runtime.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/13-arrays
go test ./curriculum/modules/03-programming-fundamentals/lessons/13-arrays
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Confusing arrays with slices - [3]int and []int are different types.
- Assuming arrays are reference types - arrays are values, passing [1024]int copies 8KB.
- Thinking you can resize an array - size is part of type, fixed at compile time.

## In Production

Arrays underpin slices - every slice has backing array. Cryptographic hashes produce [32]byte. Fixed-size network buffers use arrays.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-14`.
