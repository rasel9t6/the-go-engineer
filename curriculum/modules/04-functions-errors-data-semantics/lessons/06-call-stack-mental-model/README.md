# Call stack mental model

## Mission

Understand and apply Call stack mental model in the context of professional Go software engineering.

## Prerequisites

- core-04-05

## Mental Model

The call stack is like a stack of sticky notes. Each function call puts a new note on top. The note has the function's local variables and a 'return here' address. When the function finishes, the note is removed and execution follows the return address.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go does not use the hardware stack for goroutines -- each goroutine has its own stack allocated on the heap. The stack is dynamically grown: when a function call would exceed the current stack, the runtime allocates a larger stack (2x), copies all frames, and adjusts pointers.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/06-call-stack-mental-model
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/06-call-stack-mental-model
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Thinking the call stack is infinite -- deep recursion causes stack overflow.
- Confusing the stack (function calls) with the heap (dynamic allocation).
- Expecting goroutines to share the same call stack -- each goroutine has its own.
- Believing local variables are always on the stack -- they escape to the heap when the compiler decides they must.

## In Production

Every time you see a panic stack trace in Go, you are reading the call stack. Debuggers display the call stack to show how execution reached the current line. Profilers sample the call stack to determine where CPU time is spent.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-07`.
