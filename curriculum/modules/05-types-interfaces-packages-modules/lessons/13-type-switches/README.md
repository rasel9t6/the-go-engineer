# Type switches

## Mission

Understand and apply Type switches in the context of professional Go software engineering.

## Prerequisites

- core-05-12

## Mental Model

A type switch is like a package sorting machine. Packages arrive on a conveyor belt (interface). The machine checks each label (dynamic type). Books go to books bin, Electronics to electronics bin, unknown to default.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The compiler generates type assertions wrapped in a switch. It can optimize using the type pointer from the interface's itab. For interface{}, the compiler may use a hash-based jump table for O(1) dispatch.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/13-type-switches
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/13-type-switches
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Forgetting the default case in a type switch -- unexpected types pass silently.
- Using `x := x.(type)` but then using x as the wrong type in a case arm.
- Thinking type switches only work with `interface{}` -- they work with any interface.
- Expecting fallthrough -- Go type switches don't fall through.

## In Production

Production Go uses type switches for: JSON unmarshaling (different JSON types), error classification (net.Error, *os.PathError, custom errors), logging argument formatting, template function arguments.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-14`.
