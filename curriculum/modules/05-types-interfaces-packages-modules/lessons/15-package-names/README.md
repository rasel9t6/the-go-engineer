# Package names

## Mission

Understand and apply Package names in the context of professional Go software engineering.

## Prerequisites

- core-05-14

## Mental Model

A package name is like a department name. You don't say 'Common Things Department' -- Engineering, Sales, Marketing. Each has clear responsibility.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Package name must match across files in a directory. The compiler resolves import path to directory, reads .go files, verifies same package name. Package name is the prefix for exported identifiers.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/15-package-names
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/15-package-names
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using generic names like `common`, `utils`, `helpers` -- dumping grounds.
- Naming package same as primary export (package `config` with type `Config`) -- stutter.
- Using underscores in package names -- Go convention is lowercase single-word.
- Mixing plural and singular inconsistently.

## In Production

Production Go follows: `api`, `store`, `worker`, `server`, `config`, `middleware`. Kubernetes uses `pkg/api`, `pkg/client`. Docker uses `cli`, `client`, `daemon`.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-16`.
