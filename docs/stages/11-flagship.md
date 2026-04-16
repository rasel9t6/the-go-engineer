# 10 Flagship Project

## Purpose

`10 Flagship Project` is the long-running proof surface for the curriculum.

## Who This Is For

- learners who want one product spine that integrates many stages together
- developers who want a stronger portfolio-quality proof path inside the repo

## Mental Model

A flagship project is not a random extra app.
It is the place where backend, concurrency, quality, architecture, operations, and judgment start
to interact inside one evolving system.

## Why This Stage Exists

This stage exists so the curriculum has one long-running proof surface instead of only isolated
milestones.

The goal is to give learners a system they can extend, review, debug, and improve across multiple
engineering dimensions without losing the thread.

## What You Should Learn Here

- staged feature growth instead of one giant final dump
- architectural and operational trade-offs inside a single system
- checkpoint-driven improvement
- cross-stage integration from backend through production

## Stage Shape

This stage currently has one primary flagship seed plus one supporting pattern source:

1. `enterprise-capstone`
   - the main flagship project seed on `main`
2. milestone project patterns already used across the alpha curriculum
   - supporting examples for how checkpoints, scoped increments, and proof surfaces can work

That means the stage is anchored in one real project spine while still learning from the smaller
project patterns already present elsewhere in the repo.

## Current Source Content

- [14-application-architecture/enterprise-capstone](../../14-application-architecture/enterprise-capstone/)
- milestone project patterns already used across the alpha curriculum

## Stage Support Docs

Use these support docs when you want the beta-stage view of the flagship path:

- [Flagship Project support index](./flagship-project/README.md)
- [Stage map](./flagship-project/stage-map.md)
- [Checkpoint guidance](./flagship-project/checkpoint-guidance.md)

## Where This Stage Starts

Start with [14-application-architecture/enterprise-capstone](../../14-application-architecture/enterprise-capstone/).

That is the current flagship seed on `main`.
The other milestone project patterns support the stage, but the capstone project is the primary
spine.

## Recommended Order

Use this order for the current beta-facing shell:

1. read and run the enterprise capstone so the system shape is concrete
2. identify the first checkpoint boundary before making large changes
3. extend the project in staged increments instead of treating it as one giant final dump
4. use smaller milestone project patterns as reference when you need scope discipline

## Path Guidance

### Full Path

Enter this stage after the earlier build/test/structure/operate stages are already meaningful.
The flagship should integrate those habits, not replace them.

### Bridge Path

If you already build real services, you can move into this stage earlier, but only if you can
already explain:

- backend boundaries
- cancellation and concurrency safety
- testing and profiling habits
- operational concerns like logs and shutdown

### Targeted Path

If your immediate goal is portfolio-quality proof work, use this stage as the place where the
earlier stage habits are recombined into one system.

## Checkpoint Backbone

The current public checkpoint model is:

- foundation checkpoint: run the system, understand the moving parts, and explain the project shape
- architecture checkpoint: explain the major package and service boundaries
- operations checkpoint: explain runtime, deployment, and lifecycle behavior
- iteration checkpoint: extend the system deliberately without collapsing the design

## Finish This Stage When

- you can grow one system across multiple stages without losing structure
- you can explain the design and operations decisions in the project
- you can improve the system through iteration instead of only feature addition
- you can use the project as real proof of engineering growth

More concretely:

- you can explain why the project is split the way it is
- you can choose the next checkpoint based on real risk and learning value
- you can connect implementation choices back to testing, operations, and maintainability
- you can use the flagship as proof of engineering growth instead of as a random finished demo

## Next Stage

Move to [11 Code Generation](./11-code-generation.md) once you already understand the kinds of
systems and code you want generation tools to accelerate.
