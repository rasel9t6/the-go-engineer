# Error wrapping

## Mission

Understand and apply Error wrapping in the context of professional Go software engineering.

## Prerequisites

- core-04-09

## Mental Model

Error wrapping is like putting an error in an envelope with a note. The envelope has your message, and the original error is inside. `errors.Is` opens all envelopes to find the original.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

`fmt.Errorf` with `%w` creates a `*fmt.wrapError` struct storing the format string and wrapped error. Its `Unwrap()` returns the inner error. `errors.Is` walks the chain by calling `Unwrap()` repeatedly.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/10-error-wrapping
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/10-error-wrapping
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using `fmt.Errorf` with `%w` but not importing the `errors` package.
- Wrapping errors in a loop -- each wrap adds another layer, making the chain harder to read.
- Wrapping an error with `fmt.Errorf` but not using `%w` -- the wrapping is lost and `errors.Is` won't match.
- Wrapping errors unnecessarily -- not every error needs to be wrapped.

## In Production

Production Go services wrap errors at every I/O boundary: `db.Query` wraps with the query name, `http.Handler` wraps with the endpoint, `storage.Get` wraps with the key.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-11`.
