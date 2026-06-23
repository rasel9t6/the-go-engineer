# Function basics

## Mission

Understand and apply Function basics in the context of professional Go software engineering.

## Prerequisites

- core-03-23

## Mental Model

A function is a reusable block of code with a name, optional inputs (parameters), and optional outputs (return values). The function signature is the contract — it tells callers what to provide and what to expect back.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

In Go's compiled output, each function is a block of machine code with a known entry point. The calling convention passes arguments on the stack (or in registers on newer architectures). The stack frame holds the return address, parameters, local variables, and space for return values. Go's stack is dynamically grown — if a function call would overflow the current stack, the runtime copies the stack to a larger block, updating pointers.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/01-function-basics
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/01-function-basics
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Declaring a function inside another function and expecting it to be accessible outside — Go only supports package-level and method-level functions, not nested named functions.
- Forgetting to call the function — writing `myFunc` instead of `myFunc()` compiles but does nothing, leading to confusing no-op behavior.
- Omitting return statement in a function that declares return types — causes a compile error, but the error message may not clearly point to the missing return.

## In Production

Every Go package uses functions. HTTP handlers are functions, database query wrappers are functions, middleware chains compose functions. In production Go code, functions are the boundary between packages — a package's exported functions define its API surface.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-02`.
