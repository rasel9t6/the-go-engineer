# Beginner pointer mistakes

## Mission

Understand and apply Beginner pointer mistakes in the context of professional Go software engineering.

## Prerequisites

- core-03-22

## Mental Model

Beginner pointer mistakes share pattern: assuming pointer is valid when nil, or assuming Go magic. Fix: verify initialization before dereference.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Every Go program using pointers can crash from nil dereference. Runtime check inserted by compiler for every dereference. Loop capture bug fixed in Go 1.22 (per-iteration variable).

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/23-beginner-pointer-mistakes
go test ./curriculum/modules/03-programming-fundamentals/lessons/23-beginner-pointer-mistakes
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assigning to *p when p is nil panics - always check p != nil first.
- Using & before literal like &42 - only variables have addresses. Use new(int).
- Returning pointer to loop variable and using after loop - all captured pointers point to same final value.

## In Production

Every code review catches pointer mistakes. Service crashes from nil pointer panics are common. Knowing common errors eliminates runtime failures.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `module checkpoint`.
