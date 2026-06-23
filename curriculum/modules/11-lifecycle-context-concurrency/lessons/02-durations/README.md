# Durations

## Mission

Understand and apply Durations in the context of professional Go software engineering.

## Prerequisites

- core-11-01

## Mental Model

time.Duration is a count of nanoseconds with a nice String() method. It is Go's answer to 'how long' — a first-class numeric type with units baked into the type system.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

time.Duration is type Duration int64 with nanosecond base. Constants are integer multipliers. ParseDuration splits on letter boundaries and sums nanosecond values.

## Run Instructions

```bash
go run ./curriculum/modules/11-lifecycle-context-concurrency/lessons/02-durations
go test ./curriculum/modules/11-lifecycle-context-concurrency/lessons/02-durations
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using int64 for duration math instead of time.Duration — loses type safety.
- Confusing nanoseconds with milliseconds — using 1000 when you mean 1000 * time.Millisecond.
- Adding time.Duration directly to time.Time without using .Add() — the compiler correctly prevents this.

## In Production

Durations appear in every Go service: HTTP client timeouts, DB pool settings, cache TTLs, retry intervals, health check frequencies.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-11-03`.
