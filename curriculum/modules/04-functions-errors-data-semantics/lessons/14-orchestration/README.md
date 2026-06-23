# Orchestration

## Mission

Understand and apply Orchestration in the context of professional Go software engineering.

## Prerequisites

- core-04-13

## Mental Model

An orchestrator is like a stage manager in a theater production. The manager doesn't act or build sets -- they coordinate: cue lights, cue sound, cue actor. If the lights fail, the manager stops the show.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The orchestrator pattern maps to Go's sequential execution model. Each function call is synchronous -- the next step runs only after the previous step returns. For parallel steps, use sync.WaitGroup + goroutines.

## Run Instructions

```bash
go run ./curriculum/modules/04-functions-errors-data-semantics/lessons/14-orchestration
go test ./curriculum/modules/04-functions-errors-data-semantics/lessons/14-orchestration
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Making a function do too many things -- violating single responsibility.
- Mixing business logic, I/O, and error handling with no structure in an orchestrator.
- Skipping error handling in an orchestrator and letting a failed step corrupt the result.
- Writing orchestrators that are untestable because they hardcode dependencies.

## In Production

Every HTTP handler in a production Go service is an orchestrator. Service layer orchestrators coordinate: load user, check permissions, execute operation, audit log, return result. Clean Architecture uses orchestrators at the use-case layer.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-04-15`.
