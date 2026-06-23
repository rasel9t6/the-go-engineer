# Breakpoints

## Mission

Understand and apply Breakpoints in the context of professional Go software engineering.

## Prerequisites

- core-06-14

## Mental Model

Breakpoints are like checkpoints in a video game. You place them at critical locations. When the program reaches a checkpoint, the game pauses and you can look around (inspect state). Then you resume and continue to the next checkpoint.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Delve sets breakpoints by replacing the instruction at the target address with a software interrupt instruction (int3 on x86). When the CPU executes this instruction, the target process stops, and Delve catches the signal.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/15-breakpoints
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/15-breakpoints
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Setting breakpoints on every line instead of strategic locations.
- Not removing debug print statements after finding the bug.
- Setting breakpoints in optimized code -- may not trigger correctly.
- Forgetting that breakpoints persist across debugging sessions if not removed.

## In Production

Breakpoints are the most-used Delve feature. Developers set breakpoints at the start of suspicious functions, at key decision points, and inside loops. They're the first step in any debugging session.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-16`.
