# Comma-ok idiom

## Mission

Understand and apply Comma-ok idiom in the context of professional Go software engineering.

## Prerequisites

- core-03-17

## Mental Model

Comma-ok is two-part answer: 'here's value, and yes/no - it worked.' Like vending machine: get snack and 'yes, here' or zero-snack and 'slot empty.'

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Comma-ok is built into language for three operations. Map: hash lookup + found boolean. Type assertion: dynamic type comparison + true/false. Channel: read from buffer + true, or zero+false if closed and empty.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/18-comma-ok-idiom
go test ./curriculum/modules/03-programming-fundamentals/lessons/18-comma-ok-idiom
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Forgetting comma-ok returns (value, bool) - using v := m[key] not knowing if key existed.
- Dropping second return on type assertion - x.(T) panics if wrong; x, ok := x.(T) is safe.
- Using comma-ok on channel without select - v, ok := <-ch returns false when closed and empty.

## In Production

Comma-ok is standard for safe type extraction, safe map access, safe channel reading. Every Go codebase relies on it.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-19`.
