# Slices

## Mission

Understand and apply Slices in the context of professional Go software engineering.

## Prerequisites

- core-03-13

## Mental Model

A slice is a window into an array. Window has position (ptr), width (len), max width (cap). Slide window (reslice), but underlying array fixed. Wider window (append) needs new array if at max.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Slice is struct{ ptr *T, len int, cap int } - 24 bytes. append checks len < cap; if true, sets s[len]=v and increments len; if false, calls runtime.growslice allocating new array.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/14-slices
go test ./curriculum/modules/03-programming-fundamentals/lessons/14-slices
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Treating nil and empty slices differently - len(nil) == 0, append works on nil.
- Assuming slicing [:j] creates new copy - shares backing array with original.
- Using append without saving return value: append(s, v) discards result.

## In Production

Every DB result, file read, HTTP response body is a slice. JSON, string processing, protobuf all produce slices.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-15`.
