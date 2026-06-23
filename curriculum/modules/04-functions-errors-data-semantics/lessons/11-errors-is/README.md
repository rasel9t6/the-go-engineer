# errors.Is

## Mission

Understand and apply errors.Is in the context of professional Go software engineering.

## Prerequisites

- core-04-10

## Mental Model

`errors.Is` is like a detective following a chain of envelopes. Each envelope has a `Unwrap()` that opens to reveal the next. The detective checks each envelope until finding the target or running out of envelopes.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

`errors.Is` compares using `==`, then checks for an `Is(err error) bool` method, then calls `Unwrap()` to walk the chain. The chain walk is depth-first.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/11-errors-is
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/11-errors-is
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using `errors.Is` with a non-sentinel error value.
- Calling `errors.Is(err, target)` with an error not created with `errors.New` or `fmt.Errorf`.
- Checking `err == sentinel` instead of `errors.Is(err, sentinel)` when err might be wrapped.
- Using `errors.Is` in an if statement but continuing execution without handling the case.

## In Production

`errors.Is(err, fs.ErrNotExist)` is used in every Go file-reading codebase. `errors.Is(err, io.EOF)` in every parsing library. `errors.Is(err, context.Canceled)` in every concurrent Go service.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-12`.
