# Inspecting variables

## Mission

Understand and apply Inspecting variables in the context of professional Go software engineering.

## Prerequisites

- core-06-16

## Mental Model

Inspecting variables is like taking inventory in a store. You list all items (locals), check specific items (print), understand what type of item it is (whatis), and look inside containers (dereference pointers, inspect struct fields).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Delve reads the target process's memory through the OS debugging interface. It parses the DWARF debug information to know variable types and locations (register, stack, heap). It formats the value according to Go's formatting rules.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/17-inspecting-variables
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/17-inspecting-variables
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Forgetting to print pointer values and printing the address instead of the value.
- Not inspecting variables before the buggy line -- the bug may be in state setup.
- Only printing scalars and not inspecting struct, slice, map contents.
- Not using `whatis` to understand a variable's type when confused.

## In Production

Variable inspection is the most frequent debugging action. Developers check function parameters at entry, intermediate values during computation, and return values before the function returns to verify correctness.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-18`.
