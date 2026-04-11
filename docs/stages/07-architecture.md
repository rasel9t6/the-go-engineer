# 7 Architecture

## Purpose

`7 Architecture` teaches how systems are shaped at the package and service boundary level.

## Who This Is For

- learners who already build, test, and profile non-trivial systems
- developers who want stronger structure and boundary reasoning

## Mental Model

Architecture is not diagram theater.
It is the discipline of choosing boundaries that keep code understandable, testable, and able to
change safely.

## Why This Stage Exists

This stage exists to move learners from "I can build features" to "I can shape systems
deliberately."

The goal is to teach boundary judgment: package seams, service contracts, and structure inside a
larger codebase.

## What You Should Learn Here

- package boundary design
- service structure and interface seams
- architectural trade-offs in real systems
- boundary choices inside a capstone-sized codebase

## Stage Shape

This stage currently has one live public path plus two important architecture-owned reference
surfaces:

1. `package-design`
   - the current live beta path for package naming, visibility, and project layout
2. `grpc`
   - an architecture reference surface for schema-first service boundaries and RPC-style contracts
3. architectural slices of `enterprise-capstone`
   - a larger-system seed that shows how package and service boundaries hold together under more
     pressure

That means the stage is honest about where learners should complete proof work now while still
surfacing the broader architecture material already in the repo.

## Current Source Content

- [14-application-architecture/package-design](../../14-application-architecture/package-design/)
- [14-application-architecture/grpc](../../14-application-architecture/grpc/)
- architectural slices of [14-application-architecture/enterprise-capstone](../../14-application-architecture/enterprise-capstone/)

## Stage Support Docs

Use these support docs when you want the beta-stage view without digging through all of Section
`14`:

- [Architecture support index](./architecture/README.md)
- [Stage map](./architecture/stage-map.md)
- [Milestone guidance](./architecture/milestone-guidance.md)

## Where This Stage Starts

Start with [14-application-architecture/package-design](../../14-application-architecture/package-design/).

That is the current live beta entry because package boundaries are the foundation for everything
else in this stage.
The `grpc` and capstone architecture surfaces become more useful once that boundary vocabulary is
already clear.

## Recommended Order

Use this order for the current beta-facing path:

1. complete the package-design track from `PD.1` through `PD.3`
2. use `grpc` as the next architecture reference surface for service contracts and RPC boundaries
3. use the architectural slices of `enterprise-capstone` once you want a larger system to read and
   reason about

## Path Guidance

### Full Path

Complete the package-design track first, then move into the reference surfaces.
This keeps the stage grounded in boundary discipline instead of jumping straight into framework or
capstone complexity.

### Bridge Path

You can move faster if package design, interfaces, and service boundaries already feel somewhat
familiar, but do not skip:

- `PD.1`
- `PD.2`
- `PD.3`

Those are the main proof surfaces that show you can reason about architectural boundaries, not just
consume architecture vocabulary.

### Targeted Path

If you enter this stage with a narrow goal:

- start with `package-design` if your gap is package boundaries and structure
- use `grpc` if your gap is service contracts and RPC architecture
- use `enterprise-capstone` architecture slices if your gap is reading larger systems

## Stage Milestones

The current live milestone backbone is:

- `PD.3` project layout

The `grpc` and `enterprise-capstone` surfaces are important architecture seeds, but they are not
yet promoted as the main public beta proof path.

## Finish This Stage When

- you can explain why a system is split the way it is
- you can spot weak package boundaries and suggest better ones
- you can discuss service seams and ownership with evidence
- you can read a larger codebase without losing the architectural thread

More concretely:

- you can explain why package boundaries should reflect domain responsibilities instead of folder
  fashion
- you can discuss service contracts and interface seams without hand waving
- you understand which Section `14` surfaces belong to Architecture versus Production Engineering
  versus Flagship Project
- you can complete `PD.3` and use the reference surfaces to extend your architectural judgment

## Next Stage

Move to [8 Production Engineering](./08-production-engineering.md).
