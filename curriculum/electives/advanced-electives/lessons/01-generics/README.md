# Generics

## Mission

Understand and apply generics in the context of professional Go software engineering.

## Prerequisites

- module-05

## Mental Model

Generics allow writing type-safe, reusable code without sacrificing compile-time type checking. The type parameter acts as a stand-in that the compiler replaces with the actual type argument during compilation.

## Visual Model

```text
generic function -> concrete type -> monomorphized function -> call site
```

## Machine View

Generics are compile-time templates. When you write a generic function, the compiler creates a concrete copy of the function for each distinct set of type arguments at each call site. The generic function itself never exists at runtime — only the monomorphized concrete functions do.

## Run Instructions

```bash
go run ./curriculum/electives/advanced-electives/lessons/01-generics
go test ./curriculum/electives/advanced-electives/lessons/01-generics
```

## Try It

Write a generic function that works with any numeric type.

## In Production

Generics reduce boilerplate in data structure implementations, collection utilities, and cross-cutting concerns like retry or caching logic.

## Thinking Questions

- When should you use generics vs interfaces?
- What are the performance characteristics of monomorphized generics?

## Next Step

Complex generic constraints
