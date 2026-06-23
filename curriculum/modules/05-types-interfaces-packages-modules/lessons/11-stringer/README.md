# Stringer

## Mission

Understand and apply Stringer in the context of professional Go software engineering.

## Prerequisites

- core-05-10

## Mental Model

Stringer is the 'make this type printable' interface. It's a name tag: 'when asked what I am, tell them this.' Without it, Go prints raw struct fields. With it, you control the representation.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The `fmt` package does a type assertion to check for `fmt.Stringer`. If yes, calls `String()`. If no, falls to default formatting. The check is O(1) -- just an itab comparison.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/11-stringer
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/11-stringer
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Infinite recursion in String() -- using `%v` on the same type inside String().
- Creating String() that modifies the receiver -- Stringer should be read-only.
- Forgetting that String() must be defined on the type, not a pointer (unless both needed).
- Overriding String() but not considering pointer vs value receiver implications for fmt.

## In Production

Every custom Go type that is logged should implement Stringer. URL, IP, MAC address types implement Stringer for display. Error types implement Error() for human-readable messages.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-12`.
