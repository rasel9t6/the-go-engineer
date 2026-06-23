# Refactoring safely

## Mission

Understand and apply Refactoring safely in the context of professional Go software engineering.

## Prerequisites

- core-06-18

## Mental Model

Refactoring is like renovating a house. You don't tear down all walls at once. You renovate one room, test the structure, then move to the next room. Tests are your building inspector -- they tell you if the renovation weakened the structure.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's fast compilation enables rapid refactoring feedback. `go test` compiles and tests in seconds. `gofmt` ensures consistent formatting. `go vet` detects type and logic issues. `go doc` helps understand existing interfaces.

## Run Instructions

```bash
go run ./curriculum/modules/06-testing-debugging-refactoring/lessons/19-refactoring-safely
go test ./curriculum/modules/06-testing-debugging-refactoring/lessons/19-refactoring-safely
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Refactoring without tests first -- no safety net.
- Changing too much at once instead of small, verifiable steps.
- Refactoring and adding features simultaneously.
- Not running the full test suite after each refactoring step.

## In Production

Production codebases are constantly refactored: renaming for clarity, extracting functions for reuse, simplifying complex conditionals, and restructuring packages. Safe refactoring is an essential engineering skill.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-06-20`.
