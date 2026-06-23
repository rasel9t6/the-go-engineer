# Type assertions

## Mission

Understand and apply Type assertions in the context of professional Go software engineering.

## Prerequisites

- core-05-11

## Mental Model

A type assertion asks 'are you actually a specific brand?' The generic USB device (interface) is asked: are you a printer? If yes, printer-specific features are accessible.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The interface value stores an itab pointer (or eface for empty interfaces). The assertion compares the type pointer against the target. If match, the concrete value is extracted from the data pointer.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/12-type-assertions
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/12-type-assertions
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using type assertion on a nil interface value -- panics (use comma-ok form).
- Forgetting the comma-ok idiom: single-value form panics on failure.
- Asserting to the wrong type -- must match exactly, not just implement the same interface.
- Using assertions instead of interface methods -- if asserting frequently, reconsider interface design.

## In Production

Production Go uses type assertions for: checking if ResponseWriter supports Hijacker, if Reader supports Seeker, if error is a custom type (via errors.As).

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-13`.
