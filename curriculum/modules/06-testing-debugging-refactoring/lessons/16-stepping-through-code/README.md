# Stepping through code

## Mission

Understand and apply Stepping through code in the context of professional Go software engineering.

## Prerequisites

- core-06-15

## Mental Model

Stepping through code is like reading a book one word at a time. `next` reads the current sentence and moves to the next. `step` follows a footnote and comes back. `stepout` skips back to the main text from a footnote.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

step sets a breakpoint on the next instruction (next line), step sets a breakpoint at the target function entry, stepout sets a breakpoint at the return address. Each is a temporary breakpoint that's removed after it triggers.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/16-stepping-through-code
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/16-stepping-through-code
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Stepping into standard library functions when the bug is in application code.
- Using `next` when `step` is needed -- confusing the two.
- Not stepping through the code systematically.
- Stepping too quickly and missing the critical state change.

## In Production

Stepping is used to trace the exact execution path through unfamiliar code, to observe state changes in detail, and to understand complex algorithms one step at a time.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-17`.
