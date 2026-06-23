# Debugging mindset

## Mission

Understand and apply Debugging mindset in the context of professional Go software engineering.

## Prerequisites

- core-06-10

## Mental Model

Debugging is like being a detective at a crime scene. You gather evidence (logs, stack traces, variable values), form theories (hypotheses), test each theory, and eliminate possibilities until only the cause remains. Random guessing is not detective work.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go programs produce deterministic output for given inputs -- this makes reproduction reliable. Panics produce stack traces with file and line numbers. The race detector catches concurrency bugs. Delve provides step-by-step execution.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/11-debugging-mindset
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/11-debugging-mindset
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Treating debugging as guesswork instead of systematic investigation.
- Making random changes to see if the bug goes away -- cargo cult debugging.
- Not reproducing the bug consistently before trying to fix it.
- Skipping the step of understanding the root cause before applying a fix.

## In Production

Professional debugging follows scientific method: hypothesis, prediction, experiment, analysis. Senior engineers spend significant time debugging and use systematic approaches rather than random changes.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-12`.
