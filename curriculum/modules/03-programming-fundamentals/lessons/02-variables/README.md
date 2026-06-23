# Variables

## Mission

Understand and apply Variables in the context of professional Go software engineering.

## Prerequisites

- core-03-01

## Mental Model

A variable is a named storage location with a fixed type. A labeled box: label = name, box size = type, contents = value. Box cannot change size (type), only contents.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

var x int reserves space for an int-sized value. The variable is a named memory address. Go has no uninitialized variables - every declaration produces a valid zero value.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/02-variables
go test ./curriculum/modules/03-programming-fundamentals/lessons/02-variables
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using := at package scope - short var declaration only works inside function bodies.
- Shadowing a variable in inner scope expecting the outer value to change.
- Declaring var x int then assigning x = "hello" - Go catches type mismatch at compile time.

## In Production

Variables appear in every Go program - from err in if err != nil to config structs. Solid understanding prevents entire categories of bugs.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-03`.
