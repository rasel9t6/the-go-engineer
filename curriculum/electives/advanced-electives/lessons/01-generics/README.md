# Generics

## Mission

Understand and apply Generics in the context of professional Go software engineering.

## Prerequisites

- core-16-15

## Mental Model

Generics are compile-time templates. When you write a generic function, the compiler creates a concrete copy of the function for each distinct set of type arguments at each call site. The generic function itself never exists at runtime — only the monomorphized concrete functions do. The type parameter is a stand-in that the compiler replaces with the actual type argument during compilation.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go generics use monomorphization: the compiler generates a separate concrete function for each unique set of type arguments. This means generics have zero runtime cost — the machine code for a generic function is identical to hand-written concrete code. The tradeoff is binary size and compile time: each distinct instantiation adds code to the binary. The compiler also checks that the type argument satisfies the constraint at compile time, preventing any runtime type assertion failures.

## Run Instructions

```bash
Read the lesson and complete the practice task.
No automated test is required for this lesson.
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using interface{} instead of a type parameter — loses compile-time type safety and requires runtime type assertions that can panic.
- Declaring type parameters that are never used by the function body — adds complexity without benefit.
- Using too many type parameters — each type parameter increases compile time and makes the function signature harder to read.
- Writing constraint interfaces that are too restrictive — prevents valid type arguments from being used, and the function is less reusable than necessary.
- Forgetting that methods cannot have their own type parameters in Go 1.18-1.23 — standalone generic functions are needed instead.

## In Production

Generics are used in Go libraries for type-safe data structures (slices, maps, sets, heaps), algorithm functions (sort, map, reduce, filter), and serialization libraries. The standard library's slices.Sort, slices.Compact, and maps.Clone are built on generics. Production Go code uses generics for type-safe repository patterns, event handlers, and middleware that operate on different types without type assertions.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-02`.
