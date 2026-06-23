# Type conversions

## Mission

Understand and apply Type conversions in the context of professional Go software engineering.

## Prerequisites

- core-03-04

## Mental Model

A type conversion is an explicit instruction to reinterpret or transform a value from one type to another. Like translation - specify target language, accept possible meaning loss (truncation).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go has no implicit numeric conversions. Compiler checks convertibility: numeric types, strings <-> []byte/rune, pointer/interface conversions. Runtime handles bit manipulation. Key difference from assertions: conversion works on concrete types at compile time.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/05-type-conversions
go test ./curriculum/modules/03-programming-fundamentals/lessons/05-type-conversions
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming Go has implicit numeric conversion - even int to int64 requires explicit T().
- Confusing type conversion (T(x)) with type assertion (x.(T)) - conversion on concrete, assertion on interface.
- Using strconv.Atoi instead of understanding string to int requires runtime parsing, not conversion.

## In Production

API handlers convert strings to numbers, DB drivers convert Go types to SQL types, JSON serialization converts structs to bytes. Every production Go program performs type conversions.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-06`.
