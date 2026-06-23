# Values and expressions

## Mission

Understand and apply Values and expressions in the context of professional Go software engineering.

## Prerequisites

- core-02-14

## Mental Model

An expression is a recipe producing a value. Ingredients (operands) and tools (operators) combine in a precise order. If you know types and operators, you can trace every result.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go expressions are parsed into an AST. The type checker validates operator types match. At runtime, operands evaluate left-to-right, then operators apply by precedence. No hidden coercion exists.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/01-values-and-expressions
go test ./curriculum/modules/03-programming-fundamentals/lessons/01-values-and-expressions
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Confusing expression evaluation order - assuming left-to-right when Go evaluates operands as per spec.
- Writing 42 + "hello" expecting implicit coercion - Go requires explicit type conversion.
- Forgetting untyped constants have a default type that affects valid operations.

## In Production

Every Go program is a tree of expressions. From arithmetic in CLI tools to boolean logic in authorization, understanding expression evaluation is foundational.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-02`.
