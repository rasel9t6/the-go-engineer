# Why architecture exists

## Mission

Understand and apply Why architecture exists in the context of professional Go software engineering.

## Prerequisites

- core-13-08

## Mental Model

Architecture is not a blueprint drawn before construction. It is the set of decisions that are expensive to change. A good architecture makes the common case easy and the rare case possible. The right architecture for a startup's 3-person MVP is a single main.go file. The right architecture for a 50-engineer org is domain-aligned packages with explicit interfaces. Architecture evolves with the team size, codebase size, and business complexity. The skill is knowing WHEN to add structure, not knowing how to draw the perfect diagram.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's package system is the architecture primitive. A package is a unit of compilation, import, and name visibility. Exported names (capitalized) form the package API. Unexported names are internal. This visibility boundary is the cheapest and most effective architecture tool — it forces a clear API surface. Imports create a directed acyclic graph of dependencies. If package A depends on package B, and B depends on A, the Go compiler rejects it. This constraint prevents the worst kind of coupling (circular) at compile time. The go tool's cycle detection runs on every build, making circular imports a compile-time error rather than a runtime surprise.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/01-why-architecture-exists
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/01-why-architecture-exists
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Jumping straight to microservices — every new Go project starts with a gRPC service in its own repo, separate deployment, and a Kubernetes manifest. The overhead of service boundaries (network calls, serialization, deployment coordination, observability) overwhelms the team before they have product-market fit. Start with a modular monolith; extract services only when there is a proven scalability or team-boundary need.
- Using architecture astronaut patterns — implementing hexagonal architecture with 7 layers of interfaces before writing a single business function. The abstractions solve problems that do not exist yet and make the code impossible to navigate. Architecture should be refactored toward, not predicted upfront.
- Letting accidental architecture become permanent — the directory structure chosen in week 1 (pkg/, internal/, cmd/) becomes the team's mental model. When the service needs to split, the team is reluctant to restructure because 'that's how we've always done it.' Architecture must be treated as malleable: the right structure for 1K LOC is wrong for 100K LOC.

## In Production

Go projects that neglect architecture become unmaintainable at ~50K LOC (the 'Go monolith wall'). Projects like Kubernetes (2M+ LOC Go), Docker, and CockroachDB all went through major architecture refactors as they grew. The pattern is consistent: start with a flat structure, extract domain packages when circular imports appear, extract services when team boundaries form. Architecture in production Go is a continuous evolution, not a one-time design.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-02`.
