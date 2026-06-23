# If and else

## Mission

Understand and apply If and else in the context of professional Go software engineering.

## Prerequisites

- core-03-08

## Mental Model

if is a fork in the road. Init statement (if x := compute(); x > 0) is a prep table - compute, check, both in same step before the fork.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's if doesn't need parentheses but requires braces. Init statement is unique Go feature - declares vars visible only within if/else block. Compiler generates conditional jump instruction.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/09-if-and-else
go test ./curriculum/modules/03-programming-fundamentals/lessons/09-if-and-else
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Putting curly brace on wrong line - Go requires { on same line as if/else/for.
- Else must be on same line as closing brace: } else {.
- Deep if-else chain where switch would be clearer - 3+ branches read better as switch.

## In Production

Error handling is dominated by if err != nil. Config branches, request routing, validation all use if/else.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-10`.
