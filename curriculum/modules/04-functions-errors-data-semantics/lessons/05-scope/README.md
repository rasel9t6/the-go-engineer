# Scope

## Mission

Understand and apply Scope in the context of professional Go software engineering.

## Prerequisites

- core-04-04

## Mental Model

Scope is like nested rooms: a variable declared in a room is visible in that room and any rooms within it, but not outside. Each `{` opens a new room, each `}` closes it. If an inner room has a variable with the same name as an outer room, the inner one is what you see -- shadowing.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The Go compiler tracks scope through a symbol table. Each scope is a nested entry in the table. When resolving an identifier, the compiler walks from innermost to outermost scope. The debugger (Delve) uses the same scoping rules.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/05-scope
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/05-scope
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Declaring a variable in an if statement and expecting it to be accessible after the if block.
- Shadowing a package-level variable with a local variable of the same name inside a function.
- Accessing a variable declared in a for loop outside the loop body.
- Assuming variables declared in a switch case are accessible in subsequent cases.

## In Production

In production Go code, the `if err := ...; err != nil { return }` pattern scopes error variables to exactly where they are needed. Loop counters are scoped to the loop. Every Go developer uses scope intuitively many times an hour.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-06`.
