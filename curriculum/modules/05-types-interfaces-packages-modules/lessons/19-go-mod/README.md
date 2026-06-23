# go.mod

## Mission

Understand and apply go.mod in the context of professional Go software engineering.

## Prerequisites

- core-05-18

## Mental Model

go.mod is like a shopping list for your project's dependencies. It says: 'I need these packages at these versions.' When you build, Go fetches exactly what's on the list.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

go.mod uses a line-based format with directives: `module`, `go`, `require`, `exclude`, `replace`. The `go` directive specifies the language version, not the toolchain version. `require` lists dependencies with semantic import versioning.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/19-go-mod
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/19-go-mod
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Editing go.mod by hand instead of using `go get` or `go mod edit`.
- Setting go version to latest without checking compatibility.
- Committing go.mod with untidy dependency entries.
- Using `replace` directives in go.mod unnecessarily.

## In Production

Every Go project since 1.16 has a go.mod. CI systems run `go mod tidy && go mod verify` to check consistency. go.mod is the first file a developer checks when understanding a project's dependencies.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-20`.
