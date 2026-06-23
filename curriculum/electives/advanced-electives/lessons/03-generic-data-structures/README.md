# Generic data structures

## Mission

Understand and apply Generic data structures in the context of professional Go software engineering.

## Prerequisites

- elective-02

## Mental Model

A generic data structure is a struct with a type parameter. The type parameter propagates to all methods on the struct. When you write Stack[int], the compiler monomorphizes the struct and all its methods for int. At runtime, Stack[int] and Stack[string] are completely distinct types with no shared code — each has its own set of concrete methods operating on the correct concrete type.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A generic struct is defined with type parameters: type Stack[T any] struct { items []T }. Each method on the struct uses the same type parameter: func (s *Stack[T]) Push(item T). When instantiated (Stack[int]), the compiler creates a concrete struct with all T replaced by int. This means Push on Stack[int] takes int, and Push on Stack[string] takes string. The monomorphized struct is identical to hand-written code — no boxing, no type assertions, no indirection.

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

- Using interface{} for data structure elements — loses type safety and requires type assertions on every Pop/Get call.
- Not using pointer receivers for methods that modify the structure — value receivers operate on a copy, and modifications are lost.
- Forgetting that methods on generic types cannot have additional type parameters — use standalone generic functions for operations that need a different type parameter.
- Storing different concrete instantiations in the same variable — Stack[int] and Stack[string] are different types and cannot be assigned to each other.
- Using zero-value initialization incorrectly — var s Stack[T] initializes items to nil, and Push must handle nil slices by allocating.

## In Production

Generic data structures are used in Go for type-safe collections: Set[T] for deduplication, Stack[T] for undo/redo, Queue[T] for work queues, Heap[T] for priority queues, and LRU[T] for caches. Production Go code uses generic structures to avoid interface{} boxing in hot paths where type assertions would add overhead and risk panics.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-04`.
