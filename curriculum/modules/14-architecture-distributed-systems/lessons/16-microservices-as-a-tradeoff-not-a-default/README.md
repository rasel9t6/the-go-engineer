# Microservices as a tradeoff, not a default

## Mission

Understand and apply Microservices as a tradeoff, not a default in the context of professional Go software engineering.

## Prerequisites

- core-14-15

## Mental Model

Microservices are a loan, not an investment. They give you immediate benefits (independent deployability, team autonomy) but the interest payments (operational complexity, debugging difficulty, latency, data inconsistency) are high and compound over time. The loan is worth it only if the benefits exceed the interest payments. Most teams take the loan and then are surprised by the interest. Before adopting microservices, calculate the interest: how much time will the team spend on CI/CD, observability, service discovery, and debugging network failures?

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's toolchain makes microservices practical: each service is a separate Go module with its own go.mod. Compilation takes seconds. Container images are small (scratch-based: 5-15MB per binary). The go tool does not enforce any architecture — a monolith and a set of microservices compile the same way. The difference is in the deployment topology: monolith = one binary, one deployment; microservices = many binaries, many deployments. Go's fast compilation means the build pipeline is not the bottleneck — the operational complexity of coordinating many deployments is the bottleneck.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/16-microservices-as-a-tradeoff-not-a-default
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/16-microservices-as-a-tradeoff-not-a-default
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Adopting microservices because 'everyone does it' — this is the most expensive architectural mistake a team can make. Microservices solve specific problems: independent deployability, team autonomy, polyglot persistence, and independent scaling. If the team does not have these problems, microservices add complexity without benefit.
- Believing microservices are 'faster' — development speed DECREASES initially because of: network error handling, distributed tracing, data consistency, deployment pipelines, service discovery, and API versioning. Speed increases only after the team masters the operational complexity, which takes 6-18 months.
- Using microservices to fix a team problem — a team is slow because of unclear ownership, poor communication, or lack of testing. They split into microservices hoping the architecture will fix the team problem. It does not: the same team problems manifest as integration problems between services, but now debugging is harder.

## In Production

The industry has largely rejected the 'microservices first' approach. Amazon's CTO Werner Vogels: 'It is much easier to make a monolith into microservices than to start with microservices.' Martin Fowler (ThoughtWorks): 'First make the monolith modular, then extract services only where needed.' Kelsey Hightower (Google): 'Start with a monolith. Microservices are a last resort.' The pendulum has swung back: modular monoliths are the recommended starting point for most teams. Microservices are a tool for specific problems, not a default architecture.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-17`.
