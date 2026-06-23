# GitHub Actions

## Mission

Understand and apply GitHub Actions in the context of professional Go software engineering.

## Prerequisites

- core-15-10

## Mental Model

It defines how the service is built, deployed, configured, and monitored. Getting github actions right means the difference between a service that operates smoothly and one that requires constant firefighting.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Under the hood, github actions for Go services follows the same patterns as other compiled languages but benefits from Go's small binary size, fast compilation, and zero runtime dependencies.

## Run Instructions

```bash
go run ./curriculum/modules/15-docker-cicd-deployment/lessons/11-github-actions
go test ./curriculum/modules/15-docker-cicd-deployment/lessons/11-github-actions
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Ignoring container lifecycle semantics with github actions — containers are ephemeral, not VMs. State must be externalized.
- Assuming github actions configuration from the development environment works in production unchanged — environments differ in fundamental ways.
- Not testing github actions failure scenarios — what happens when the registry is unreachable or the deploy fails mid-way.

## In Production

Production Go services use github actions for every deployment. Understanding github actions is essential for any Go engineer working on production services.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-15-12`.
