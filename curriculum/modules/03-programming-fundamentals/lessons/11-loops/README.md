# Loops

## Mission

Understand and apply Loops in the context of professional Go software engineering.

## Prerequisites

- core-03-10

## Mental Model

A loop is a repeat instruction. Init sets up counters, condition says when to stop, post updates state. for-range is like conveyor belt - each item passes by.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Compiler translates for into branch instructions. For for-range, creates hidden iterator. Special handling for strings (rune decode) and maps (random order).

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/11-loops
go test ./curriculum/modules/03-programming-fundamentals/lessons/11-loops
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Forgetting Go only has for - three forms: for init;cond;post, for cond, for range.
- Using for-range on string thinking it iterates bytes - it iterates runes, yielding byte index + rune.
- Modifying range variable thinking it affects collection - range copies each element.

## In Production

Every server runs event loop. Every pagination loops over results. Every batch job iterates. Loops are backbone of repetition.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-12`.
