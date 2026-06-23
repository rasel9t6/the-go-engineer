# Subtests

## Mission

Understand and apply Subtests in the context of professional Go software engineering.

## Prerequisites

- core-06-03

## Mental Model

Subtests are like chapters in a test book. The Test function is the book, each t.Run is a chapter. You can read (run) the whole book or jump to a specific chapter. Each chapter has its own setup and teardown.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

t.Run starts a new goroutine for each subtest. The parent test waits for all subtests to complete before returning. The testing framework tracks each subtest as a separate entry in the test result hierarchy.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/04-subtests
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/04-subtests
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Not using subtests for table-driven tests -- all cases run in one blob.
- Sharing t between parent and subtests incorrectly.
- Using `t.Fatal` in a subtest instead of `t.Error` -- Fatal stops the parent test too.
- Running subtests in parallel without understanding race conditions.

## In Production

Subtests are standard in Go testing. Table-driven tests use subtests for each case. CI systems show individual subtest pass/fail. Debuggers can target specific subtests.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-05`.
