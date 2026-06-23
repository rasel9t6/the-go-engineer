# Complex generic constraints

## Mission

Understand and apply Complex generic constraints in the context of professional Go software engineering.

## Prerequisites

- elective-01

## Mental Model

A constraint is a compile-time contract that defines which types a type parameter accepts. Simple constraints (any, comparable) accept broad categories. Complex constraints narrow the set using type unions (A | B | C), type approximations (~T accepts T and any type with underlying type T), and method requirements. The compiler enforces the constraint at every call site — if the type argument does not satisfy the constraint, the code does not compile.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's constraint interface can include type elements: type unions (int | string) and type approximations (~int). These are not methods — they are compile-time restrictions on the set of acceptable type arguments. The compiler checks constraint satisfaction during type checking, before monomorphization. A type argument satisfies a constraint if: it implements all methods in the constraint AND it is a member of the constraint's type set (for type elements). Type approximation (~T) means the type argument's underlying type must be T, not the type itself.

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

- Using any as a constraint when a more specific constraint is possible — loses compile-time type safety and allows invalid type arguments.
- Writing constraints that are too restrictive — preventing valid callers and reducing reuse, requiring duplicated functions for different types.
- Confusing interface{} (any) with a constraint interface — any accepts every type, while a constraint restricts which types are valid.
- Forgetting that ~int matches int AND any type with underlying type int (e.g., type MyInt int) — this is often unexpected.
- Using type unions with types that do not share common operations — the generic function body can only use operations common to all union members.

## In Production

Complex constraints are used in Go libraries for numeric types (constraints.Integer, constraints.Float), ordered types (constraints.Ordered), and serialization (json.Marshaler + json.Unmarshaler). The Go standard library's sort.Slice is not generic because Go 1.18 generics do not support method-level type parameters — instead, custom sort functions use constraints.Ordered for sort.SortFunc in golang.org/x/exp/slices.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-03`.
