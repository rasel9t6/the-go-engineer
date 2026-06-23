# Project layout

## Mission

Understand and apply Project layout in the context of professional Go software engineering.

## Prerequisites

- core-05-16

## Mental Model

Project layout is like a building floor plan. cmd/ is the lobby entrance. internal/ is the employee-only area. pkg/ is the public conference room. Each area has a clear purpose.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The `internal/` package restriction is enforced by the compiler: packages importing from another module's internal/ must share a common root. This enables module-level encapsulation beyond package-level export rules.

## Run Instructions

```bash
go run ./curriculum/modules/05-types-interfaces-packages-modules/labs/17-project-layout
go test ./curriculum/modules/05-types-interfaces-packages-modules/labs/17-project-layout
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Flat directory with 20+ .go files in one package -- breaks clarity.
- Mixing multiple packages in the same directory -- one package per directory.
- Putting all business logic in main.go -- main should be thin.
- Creating unnecessarily deep directory trees -- 3-4 levels is enough.

## In Production

Kubernetes, Docker, Prometheus, and most popular Go projects follow the standard Go project layout. `cmd/`, `internal/`, `pkg/` are recognizable conventions across the Go ecosystem.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-05-18`.
