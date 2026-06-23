# Modular monolith

## Mission

Understand and apply Modular monolith in the context of professional Go software engineering.

## Prerequisites

- core-14-04

## Mental Model

A modular monolith is a codebase that is structured like a set of microservices but deploys as a single binary. Each 'module' is a potential future service. The module boundary is an interface, not a network call. The difference from microservices is deployment topology (one binary vs. many), not design topology (modules with explicit interfaces).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's internal/ convention enforces module boundaries at the compiler level. A package at internal/catalog/ can only be imported by code whose root directory is an ancestor of internal/. For a project at github.com/company/service, internal/catalog/ can be imported by github.com/company/service and its subdirectories — but NOT by any external package. This makes internal/ the ideal container for modules. Within the module, the api/ subdirectory contains the exported interface types. The service implementation is in the module root or a service/ subdirectory — unexported types that cannot be imported by other modules. Go's import cycle detection (compiler error on cycles) forces the module dependency graph to be acyclic.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/05-modular-monolith
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/05-modular-monolith
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Calling a monolith 'modular' just because it has multiple Go packages — a modular monolith requires domain-aligned packages with explicit API contracts, not just splitting code into files. If every internal package depends on each other, it is a distributed monolith, not a modular one.
- Letting modules share a database schema — when modules share tables, they become coupled at the data layer. Changing the users table for the billing module breaks the notification module. A modular monolith means each module owns its data — use separate schemas or table prefixes per module.
- Using in-process function calls as the only module boundary — when modules communicate through function calls with shared in-memory state, extracting a module into a separate service later requires rewriting the communication layer. A modular monolith should use explicit interfaces as the module boundary, so extracting a module means implementing the same interface over the network.

## In Production

The modular monolith is the recommended starting architecture for most Go teams. Shopify (Ruby) famously ran a monolithic codebase for a decade before extracting services. GitHub's architecture started as a Ruby on Rails monolith. In Go, the modular monolith pattern is used by: HashiCorp's products (Terraform, Vault, Consul all started as modular monoliths), Sourcegraph (single binary with domain modules), and Mattermost (modular monolith with plugin architecture). The lesson: start modular, extract services only when there is a proven need.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-06`.
