# Boolean logic

## Mission

Understand and apply Boolean logic in the context of professional Go software engineering.

## Prerequisites

- core-03-07

## Mental Model

Boolean logic is true/false algebra. AND (&&) requires both true, OR (||) requires at least one, NOT (!) flips. Short-circuit is like lazy co-workers - first answer is enough.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Compiler generates conditional jump for each operator: && maps to JZ/JNZ, || maps to JZ/JNZ, ! maps to compare-and-invert. Side effects in right operand never execute if short-circuited.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/08-boolean-logic
go test ./curriculum/modules/03-programming-fundamentals/lessons/08-boolean-logic
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming && and || evaluate both sides - Go short-circuits when left determines result.
- Confusing bitwise (&, |) with logical (&&, ||) - & always evaluates both, && short-circuits.
- Writing complex boolean without parentheses - Go requires explicit bool values, no truthiness.

## In Production

Every if, for condition, guard expression uses boolean logic. API authorization, validation, error handling depend on clear boolean expressions.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-09`.
