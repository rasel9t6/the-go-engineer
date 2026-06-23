# Methods

## Mission

Understand and apply Methods in the context of professional Go software engineering.

## Prerequisites

- core-05-01

## Mental Model

A method is a function that belongs to a type. A Person struct with a Greet() method that says 'Hello, my name is X'. The method accesses the Person through the receiver.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Methods are syntactic sugar for functions with a receiver parameter. `func (t T) Foo()` becomes `func Foo(t T)`. The compiler determines the method set of each type.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/02-methods
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/02-methods
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Defining a method inside another method -- Go doesn't support nested methods.
- Forgetting that the method receiver is passed by value -- modifications don't affect the original.
- Declaring a method on a non-local type (from another package) -- compile error.
- Calling a method with pointer receiver on a non-addressable value -- compile error.

## In Production

Every Go type with behavior uses methods. HTTP handlers implement `ServeHTTP`. The `error` interface requires `Error() string`. Custom types use methods for JSON marshaling. Methods are standard for attaching behavior to types.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-03`.
