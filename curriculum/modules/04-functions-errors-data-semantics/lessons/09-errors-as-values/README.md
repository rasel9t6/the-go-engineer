# Errors as values

## Mission

Understand and apply Errors as values in the context of professional Go software engineering.

## Prerequisites

- core-04-08

## Mental Model

An error is a value, not a control flow mechanism. If a function can fail, it returns an error as a regular return value — the caller decides what to do with it. There is no throw, no catch, no stack unwinding. The error value is just data: it can be logged, wrapped, compared, or ignored (but should not be).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The error type in Go is a built-in interface: type error interface { Error() string }. Any type with an Error() string method satisfies it. errors.New returns a pointer to a simple struct with a string field. fmt.Errorf uses errors.New internally after formatting the string. When you return err, you return an interface value — if the underlying type is *os.PathError, the interface is non-nil even if the pointer is nil (see the nil interfaces lesson). The 'if err != nil' check is a runtime interface nil check comparing both the type and data pointers.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/09-errors-as-values
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/09-errors-as-values
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Assigning the result of a function call to a variable but checking a different variable for the error — the error is silently discarded.
- Using _ = fmt.Sprintf() to ignore an error return, believing the function never fails in practice — Sprintf documents it returns an error for a reason.
- Nesting if err != nil blocks five levels deep instead of returning early — the happy path is buried in indentation and the error paths are indistinguishable.
- Checking err != nil but doing nothing meaningful with the error — logging and continuing as if nothing happened, leaving the system in an unknown state.

## In Production

Production Go services use errors-as-values for every I/O operation, every database call, every API request, every configuration parse. Skipping an error check in production means a partial failure is invisible until it cascades into a full outage.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-10`.
