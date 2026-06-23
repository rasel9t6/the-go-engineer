# Documentation comments

## Mission

Understand and apply Documentation comments in the context of professional Go software engineering.

## Prerequisites

- core-05-22

## Mental Model

Documentation comments are like a user manual for your code. They're not implementation notes -- they tell other developers what your code does, what it expects, and what they can rely on. Good documentation makes code usable without reading the source.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's doc comment format uses a simple convention: the comment immediately preceding a top-level declaration documents that declaration. `go doc` and `pkgsite` parse these comments and display them as documentation. There's no markup language -- just plain text with code examples.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/23-documentation-comments
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/23-documentation-comments
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Writing documentation comments that just repeat the function name.
- Not documenting exported identifiers at all -- go vet will warn.
- Using `//` instead of `// ` (double-slash space) -- formatting tools expect the space.
- Writing implementation details instead of documenting the contract/documentation vs comment confusion.

## In Production

Every well-maintained Go project has doc comments. `go vet` enforces the convention. `go doc` provides command-line documentation. Go's standard library documentation is the gold standard for doc comments.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `module checkpoint`.
