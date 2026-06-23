# Dependency management

## Mission

Understand and apply Dependency management in the context of professional Go software engineering.

## Prerequisites

- core-05-20

## Mental Model

Dependency management is like planning a potluck dinner. Each guest (module) brings certain dishes (packages). The organizer (go mod tidy) makes sure there's no conflict (version mismatch) and the meal is complete (all dependencies satisfied).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Minimal version selection (MVS) works by walking the dependency graph and picking the minimum version of each module that satisfies all requirements. MVS guarantees reproducible builds -- given the same inputs, the same versions are always selected.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/lessons/21-dependency-management
go test ./curriculum/modules/05-types-interfaces-packages-modules/lessons/21-dependency-management
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Not pinning dependency versions -- using latest causes unpredictable builds.
- Adding dependencies for trivial functionality that could be written in a few lines.
- Not running `go mod tidy` before committing -- leaves unnecessary or stale entries.
- Committing vendor directory without running `go mod vendor` first.

## In Production

Production Go services manage dozens or hundreds of dependencies via go.mod. CI pipelines enforce tidy state. Dependency updates are managed carefully with semantic versioning awareness.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-22`.
