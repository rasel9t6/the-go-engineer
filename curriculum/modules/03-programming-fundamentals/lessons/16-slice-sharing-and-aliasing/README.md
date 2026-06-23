# Slice sharing and aliasing

## Mission

Understand and apply Slice sharing and aliasing in the context of professional Go software engineering.

## Prerequisites

- core-03-15

## Mental Model

Sharing is group sharing single whiteboard. Each sub-slice writes on assigned section. Writing beyond section overwrites others' work. Only independent whiteboard (copy) prevents this.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

s[low:high] creates slice header where ptr=base+low, len=high-low, cap=cap(s)-low. Append within capacity overwrites original array at indices >= high. This is intentional and predictable.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/16-slice-sharing-and-aliasing
go test ./curriculum/modules/03-programming-fundamentals/lessons/16-slice-sharing-and-aliasing
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Modifying through one sub-slice and surprised another shows the change.
- Appending to sub-slice accidentally overwriting original elements.
- Passing slice to function, appending inside, expecting caller's length to change.

## In Production

HTTP middleware chains share contexts, DB scanners share scan buffers, JSON decoders share internal buffers. Understanding aliasing prevents corruption.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-17`.
