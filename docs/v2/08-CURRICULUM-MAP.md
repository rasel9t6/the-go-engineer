# V2 Curriculum Map

Status: Draft planning map. This is a structural target, not a frozen lesson list.

## Design Direction

The first v2 draft keeps the broad topic progression of v1 because the overall sequence is already
good. The major change is not the existence of new topics. The major change is that every section
will now participate in the same training system:

- lessons
- drills
- guided exercises
- section checkpoints
- mini-projects
- capstones

## Proposed Phase Model

### Phase 1: Foundations

- Section 01: Core Foundations
- Section 02: Control Flow
- Section 03: Data Structures
- Section 04: Functions and Errors

Goal: build confidence with syntax, values, references, and Go-style error handling.

### Phase 2: Go Design

- Section 05: Types and Interfaces
- Section 06: Composition
- Section 07: Strings and Text
- Section 08: Modules and Packages

Goal: teach idiomatic Go design, abstraction, and package-level thinking.

### Phase 3: IO And Applications

- Section 09: IO and CLI
- Section 10: Web and Database

Goal: move from language fluency into runnable applications with real boundaries.

### Phase 4: Systems

- Section 11: Concurrency
- Section 12: Concurrency Patterns

Goal: teach safe, explainable concurrency before advanced system composition.

### Phase 5: Engineering Quality

- Section 13: Quality and Performance
- Section 14: Application Architecture

Goal: raise the learner from "can build" to "can ship and maintain."

### Phase 6: Professional Practice

- Section 15: Code Generation

Goal: teach automation and production-minded workflow improvements that save engineering time.

## Section Design Template

Every v2 section should define:

- mission
- learner outcomes
- prerequisite expectations
- canonical lesson list
- exercise inventory
- section checkpoint
- mini-project or capstone linkage
- navigation and documentation links

## Proposed Section Outcomes

| Section | Mission | Required Outputs |
| ------- | ------- | ---------------- |
| 01 | establish environment, syntax confidence, and the mental model of a Go program | lessons, drills, beginner checkpoint |
| 02 | teach control flow as decision-making under simple rules | lessons, guided exercise, checkpoint |
| 03 | build memory and collection intuition | lessons, guided exercise, pointer checkpoint |
| 04 | make errors and control of execution a core habit | lessons, guided exercise, first mini-project |
| 05 | teach types, methods, interfaces, and generic design | lessons, exercise, design checkpoint |
| 06 | teach composition as the default structuring tool | lessons, exercise, mini-project linkage |
| 07 | teach string, bytes, text processing, templates, and parsing | lessons, parser exercise |
| 08 | teach module boundaries, package organization, and dependency hygiene | lessons, repo-structure checkpoint |
| 09 | teach file, stream, encoding, and CLI boundaries | lessons, CLI mini-project |
| 10 | teach HTTP, persistence, handlers, and service boundaries | lessons, web/data mini-project |
| 11 | teach goroutines, channels, synchronization, and failure modes | lessons, concurrency checkpoint |
| 12 | teach production concurrency patterns and cancellation discipline | lessons, pipeline mini-project |
| 13 | teach testing, mocks, benchmarks, and profiling | lessons, quality checkpoint |
| 14 | teach packaging, logging, deployment concerns, and service architecture | lessons, architecture mini-project |
| 15 | teach generation, automation, and workflow leverage | lessons, final professional practice project |

## Section-Level Content Direction

The first planning recommendation is conservative:

- keep most section names from v1
- improve internal structure before large-scale renaming
- add missing practice layers
- reserve major section splits or merges for places where there is a clear learner benefit

## What Changes In v2

Compared with v1, the section map should gain:

- clearer output expectations per section
- visible checkpoints between lessons and large projects
- a standard mini-project ladder
- explicit links between beginner, fast-track, and deep-dive learning paths

## What Stays Stable

Compared with v1, the section map should preserve:

- the general domain ordering
- the strong later-stage focus on concurrency, testing, and architecture
- the repo's practical and build-first teaching style

## Open Decisions

These section-level decisions still need review:

- whether Section 10 should later split into separate HTTP and persistence tracks
- whether Section 15 should broaden into a larger professional tooling section
- whether any phase needs an explicit review-and-repair section
