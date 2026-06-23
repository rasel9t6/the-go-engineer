# Slice length and capacity

## Mission

Understand and apply Slice length and capacity in the context of professional Go software engineering.

## Prerequisites

- core-03-14

## Mental Model

Length = cars on lot. Capacity = total spaces. Park more cars (append) up to capacity without building new lot. Lot full = new lot, cars move.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

len and cap are compiler intrinsics. For slices, returns header fields. make allocates backing array of cap size, initializes len elements to zero. append growth factor ~2x small, ~1.25x large.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/15-slice-length-and-capacity
go test ./curriculum/modules/03-programming-fundamentals/lessons/15-slice-length-and-capacity
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Confusing len and cap - len is current elements, cap is backing array size.
- Assuming s[:cap(s)] is always safe - reslicing to cap reveals hidden backing elements.
- Using make([]T, 5) for append-oriented usage instead of make([]T, 0, 5).

## In Production

DB result batching, I/O buffer pooling, network packet processing depend on managing len/cap for allocation-free patterns.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-16`.
