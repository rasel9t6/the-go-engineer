# CF.4 Switch

## Mission

Learn how to choose between several possible paths without writing long, hard-to-read branch chains.

## Why This Lesson Exists Now

`if / else if / else` works well for a small number of branches.
When the number of cases grows, `switch` often reads more clearly.

## Production Relevance

In production Go code, switch matters because:

- **State handling**: Map request states or HTTP status codes
- **Command processing**: Handle different user commands or operations
- **Mode selection**: Choose behavior based on configuration or environment
- **Type handling**: Differentiate between categories or kinds

Real services use switch for routing, status codes, and business rule engines.

## Mental Model

Switch is like a multi-way if-else but organized as a table of cases.

Instead of checking each condition one by one, you check one value against many possibilities.

The key: only ONE case runs, and you can group multiple values in one case.

## Visual Model

```text
switch day {
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    print "weekday"
case "Saturday", "Sunday":
    print "weekend"
}
```

```text
switch {
case score >= 90:
    print "A"
case score >= 80:
    print "B"
default:
    print "other"
}
```

## Machine View

A switch evaluates one expression and compares it against each case in order.

Unlike some languages, Go does not fall through by default.
Once a case matches and runs, the switch ends.

The `default` case runs when no other case matches.

## Prerequisites

- `CF.1` if/else
- `CF.2` for basics

## Run Instructions

```bash
go run ./02-language-basics/03-control-flow/4-switch
```

## Code Walkthrough

### `day := "Monday"`

The first example starts with a simple text value.
The program wants different output depending on that value.

### `switch day { ... }`

This form compares one value against several cases.

It reads like:

- inspect `day`
- match the first fitting case
- run only that case

### `case "Saturday", "Sunday":`

One case can match more than one value.
That helps group related outcomes without duplicating code.

### `score := 82` and `switch { ... }`

This is the tagless form of `switch`.

Instead of comparing one named value, each case is a condition.
That makes it useful for ranges and guard-like logic.

## Common Mistakes

- reaching for `switch` when a single `if` would be clearer
- assuming Go falls through to the next case automatically
- forgetting that tagless `switch` cases are still checked from top to bottom

## Try It

1. Change `day` to `Sunday`.
2. Change `score` to `95`.
3. Reorder the score cases and see how that changes the result.

## Why This Matters In Real Software

`switch` is useful when code needs to react to:

- modes
- states
- commands
- categories

It often makes business rules easier to scan than long `if / else if` ladders.

## Next Step

Continue to `CF.5` pricing checkout.
