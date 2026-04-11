# 4 Backend Engineering

## Purpose

`4 Backend Engineering` is where the curriculum becomes application-shaped.

## Who This Is For

- learners who can already work with packages, files, CLI tools, and encoded data
- developers who want to move into HTTP, persistence, and service-style boundaries

## Mental Model

Backend work is request flow plus data flow.
This stage teaches how requests enter a system, how state is persisted, and how boundaries keep the
application understandable.

## Why This Stage Exists

This is where the curriculum becomes application-shaped instead of only language-shaped.

The goal is to help learners connect transport, persistence, and service boundaries without hiding
the operational realities underneath them.

## What You Should Learn Here

- HTTP client and server basics
- handlers and request/response flow
- SQL-backed persistence and repository boundaries
- transaction thinking
- migration-aware schema change workflows

## Stage Shape

This stage currently has one live public path plus additional reference surfaces:

1. `databases`
   - the current live beta path: `database/sql`, transactions, repository boundaries, and a
     runnable persistence milestone
2. `http-client`, `web-masterclass`, and `database-migrations`
   - valuable backend reference surfaces that remain available while the public beta path is still
     being promoted more selectively

That means the stage is honest about where the learner should start now while still exposing the
broader Section `10` inventory.

## Current Source Content

- [10-web-and-database/http-client](../../10-web-and-database/http-client/)
- [10-web-and-database/web-masterclass](../../10-web-and-database/web-masterclass/)
- [10-web-and-database/databases](../../10-web-and-database/databases/)
- [10-web-and-database/database-migrations](../../10-web-and-database/database-migrations/)

## Where This Stage Starts

Start with the databases track at
[10-web-and-database/databases](../../10-web-and-database/databases/).

That is the current live beta path for this stage.
The other Section `10` directories remain useful reference material, but they are not yet the main
public learner route.

## Recommended Order

Use this order for the current beta-facing path:

1. complete the databases track from `DB.1` through `DB.6`
2. use `http-client`, `database-migrations`, and `web-masterclass` as optional reference material
   while they remain outside the live beta path

## Path Guidance

### Full Path

Complete the live databases path first, then use the other Section `10` surfaces as reinforcement
instead of as your main route.

### Bridge Path

You can move faster if SQL basics, service-style code, or backend tooling already feel somewhat
familiar, but do not skip:

- `DB.1`
- `DB.3`
- `DB.5`
- `DB.6`

Those are the main proof surfaces that show you can reason about real persistence boundaries
instead of only reading backend examples.

### Targeted Path

If you enter this stage late, start with the databases track if your main gap is service-side data
flow, persistence, and transactional safety.

If your immediate goal is only reference reading around HTTP or migrations, use those folders as
supplemental material instead of treating them as the current beta learner path.

## Stage Milestones

The current live milestone backbone is:

- `DB.6` repository pattern project

That single live milestone is intentional for now.
This stage is still promoting the databases path first before the broader Section `10` inventory is
fully regrouped into beta surfaces.

## Finish This Stage When

- you can build and explain a small backend flow end to end
- you understand how handlers, data access, and transactions connect
- you can reason about schema change and persistence boundaries
- you can debug common request/data flow mistakes without panic

More concretely:

- you can explain why repository boundaries help separate application logic from SQL mechanics
- you can reason about transaction safety and context-aware queries
- you understand which parts of the current Section `10` inventory are live beta path versus legacy
  reference material
- you can complete `DB.6` without treating persistence code as hidden magic

## Next Stage

Move to [5 Concurrency System](./05-concurrency-system.md).
