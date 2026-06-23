# Delve basics

## Mission

Understand and apply Delve basics in the context of professional Go software engineering.

## Prerequisites

- core-06-13

## Mental Model

Delve is like a microscope for your running code. You can pause execution at any point (breakpoint), examine the current state (variables, goroutines, stack), and step forward one instruction at a time to observe how state changes.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Delve uses the operating system's process tracing facilities (ptrace on Linux, mach exceptions on macOS). It controls the target process, injects breakpoint instructions, and reads/writes registers and memory to inspect state.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/14-delve-basics
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/14-delve-basics
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Not installing Delve when trying to debug Go programs.
- Trying to use Delve on an optimized binary -- set -gcflags='-N -l' to disable optimization.
- Starting a debugging session without a clear goal.
- Not knowing the basic Delve commands (break, continue, print, next, step).

## In Production

Delve is used for debugging complex logic, understanding unexpected behavior, and stepping through unfamiliar code. It's essential for debugging concurrent programs where print statements are insufficient.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-15`.
