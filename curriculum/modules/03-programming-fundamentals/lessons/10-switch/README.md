# Switch

## Mission

Understand and apply Switch in the context of professional Go software engineering.

## Prerequisites

- core-03-09

## Mental Model

Switch is a multi-way branch checking a value against possibilities. Like a mail sorter - each letter goes into one bin. No hole in bin bottom (no fallthrough).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Switch compiles differently per case values. Integer constants = jump table. String constants = binary search. Type switch compares dynamic type pointer against each case type.

## Run Instructions

```bash
go run ./curriculum/modules/03-programming-fundamentals/lessons/10-switch
go test ./curriculum/modules/03-programming-fundamentals/lessons/10-switch
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assuming implicit fallthrough like C - Go only executes matching case, stops automatically.
- Forgetting that break in switch exits the switch - break is legal but not needed.
- Using switch on non-comparable types (slice, map, func).

## In Production

HTTP routing, error classification, config validation use switch. Go services rely on switch for clean multi-way branching.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-03-11`.
