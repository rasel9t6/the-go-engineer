# CF.1 If / Else

## Mission

Learn how a Go program chooses one path or another based on a condition.

This is the first control-flow building block.
Without it, a program can only run in a straight line.

## Why This Lesson Exists Now

You already know values like numbers and strings.
Now you need to decide what happens when those values mean different things.

That is what `if`, `else if`, and `else` do.

## Production Relevance

In production Go code, branching matters because:

- **Validation**: Check if input is valid before processing
- **Authorization**: Verify permissions before allowing actions
- **Error handling**: Decide what to do when something fails
- **State machines**: Transition between different program states

Real services use branching for every request validation, error response, and business rule.

## Mental Model

Branching is about questions that have limited answers.

A program asks a yes/no question (or a which-one question) and picks exactly one path to follow.

The key insight is: only ONE branch runs. The others are skipped.

## Visual Model

```text
condition: temperature > 30?
    |
    +--> true  --> print "hot"
    |
    +--> false --> print "okay"
```

```text
condition: score >= 90?
    |
    +--> true  --> "A"
    +--> false --> check score >= 80
                       |
                       +--> true --> "B"
                       +--> false --> check score >= 70
                                          |
                                          +--> true --> "C"
                                          +--> false --> "F"
```

## Machine View

When an `if` statement runs, the Go runtime evaluates the condition expression.

If the result is true, it executes the first block and skips all `else if` and `else` blocks.
If false, it checks each `else if` in order.
If none match, it runs the `else` block (if it exists).

Only the selected block's code runs. The other code never executes.

## Prerequisites

- `LB.4` application logger (or equivalent: variables, constants, iota)

## Run Instructions

```bash
go run ./02-language-basics/03-control-flow/1-if-else
```

## Code Walkthrough

### `temperature := 25`

The program starts with a value it can reason about.
That value is not special.
It is simply something the program can compare.

### `if temperature > 30 { ... } else { ... }`

This is the simplest branch shape:

- if the condition is true, run the first block
- otherwise, run the `else` block

Only one branch runs.

### `score := 85`

Now the program makes a second decision using a different kind of rule.
Instead of a yes/no check, it compares several ranges.

### `if score >= 90`, `else if score >= 80`, `else`

This chain means:

- test the first condition
- if it fails, test the next one
- keep going until one matches
- if none match, fall back to `else`

### `username := ""`

This line shows that branching is not only about numbers.
Programs also branch on text, flags, and general state.

### `if username == "" { ... }`

This branch treats the empty string as a missing value.
That idea appears everywhere in real programs:

- missing input
- incomplete configuration
- invalid request data

## Common Mistakes

- forgetting that braces are required in Go
- writing complex nested conditions before the simple case is clear
- thinking every branch chain needs `else`

## Try It

1. Change `temperature` to `35` and rerun the lesson.
2. Change `score` to `71` and see which branch runs.
3. Set `username` to your own name and inspect the final branch output.

## Why This Matters In Real Software

Branching is how software decides:

- whether input is valid
- whether access is allowed
- whether an order should continue
- whether a request should fail fast

## Next Step

Continue to `CF.2` for basics.
