# Fuzz testing

## Mission

Understand and apply Fuzz testing in the context of professional Go software engineering.

## Prerequisites

- core-06-19

## Mental Model

Fuzzing is like having a million monkeys typing random inputs to your program. The fuzzer is smarter than monkeys -- it uses coverage guidance to find inputs that explore new code paths. When a crash occurs, the input is saved for reproduction.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's fuzzer uses coverage-guided mutation: it tracks which code paths each input covers, then mutates inputs to explore uncovered paths. Crashing inputs are minimized to the smallest failing input before saving.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/20-fuzz-testing
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/20-fuzz-testing
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Writing fuzz tests for functions that don't benefit from fuzzing (simple deterministic logic).
- Not providing good seed corpus entries -- the fuzzer starts from scratch.
- Running fuzz tests without a time limit -- they run indefinitely.
- Not understanding that fuzzing finds different bugs than unit tests.

## In Production

Fuzzing is used for security-critical code: parsers, decoders, network protocol implementations. The Go standard library runs fuzz tests in CI. Fuzzing found hundreds of CVEs in various projects.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-21`.
