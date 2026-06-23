# Time basics

## Mission

Understand and apply Time basics in the context of professional Go software engineering.

## Prerequisites

- core-10-24

## Mental Model

Time in Go is a struct with two internal fields: wall clock (for calendar ops) and monotonic clock (for duration measurement). Duration operations subtract monotonic components, guaranteeing accuracy regardless of clock adjustments.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

time stores time as a 64-bit wall clock UNIX nanoseconds and a 64-bit monotonic offset from process start. The runtime initializes the monotonic clock using CLOCK_MONOTONIC.

## Run Instructions

```bash
go run ./curriculum/modules/11-lifecycle-context-concurrency/lessons/01-time-basics
go test ./curriculum/modules/11-lifecycle-context-concurrency/lessons/01-time-basics
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using time.Now() for measuring elapsed time instead of time.Since() — wall clock can jump due to NTP adjustments.
- Storing time as Unix timestamps (int64) instead of time.Time — loses timezone, monotonic clock, and formatting.
- Comparing time.Time with == instead of .Equal() — time.Time has both wall clock and monotonic components.

## In Production

Production Go services use time.Time for every timestamped operation: request logging, metrics, cache expiration, session timeouts, and scheduled jobs.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-11-02`.
