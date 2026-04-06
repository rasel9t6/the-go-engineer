# V2 Project Ladder

## Purpose

Projects should act as curriculum milestones, not as random bonus content.
The ladder exists to make project complexity grow with learner readiness.

## Project Strategy

The recommended v2 strategy is a ladder of related but independent deliverables.

That means:

- do not force a single giant monolith from Section 01 onward
- do reuse concepts and, where helpful, domains across phases
- let each project validate a distinct curriculum jump

## Project Classes

v2 should use three project classes:

### 1. Mini-Project

- section or phase scoped
- one main engineering jump
- usually solvable in a small package set

### 2. Phase Capstone

- spans several sections
- forces synthesis across domains
- should feel close to real engineering work

### 3. Final Capstone

- portfolio-quality consolidation
- uses prior sections with stronger quality, architecture, and automation expectations

## Recommended Ladder

The first v2 draft should use this ladder.

### Project 1: Foundations Mini-Project

Placement:

- after Section 04

Role:

- prove that the learner can combine variables, control flow, data structures, functions, and error
  handling into one coherent tool

Suggested archetype:

- `expense-tracker-cli`
- `task-journal-cli`

### Project 2: Design Mini-Project

Placement:

- after Section 08

Role:

- prove understanding of types, interfaces, composition, text handling, packages, and module
  boundaries

Suggested archetype:

- `catalog-domain-tool`
- `report-renderer`

### Project 3: IO And CLI Mini-Project

Placement:

- around Section 09 or immediately after

Role:

- prove comfort with files, streams, encoding, and command-line workflow design

Suggested archetype:

- `backup-and-report-cli`
- `log-ingestion-cli`

### Project 4: Web And Data Mini-Project

Placement:

- after Section 10

Role:

- prove learners can build a service boundary with persistence, routing, validation, and handler
  structure

Suggested archetype:

- `notes-service`
- `bookmark-api`

### Project 5: Concurrency Systems Project

Placement:

- after Section 12

Role:

- prove the learner can coordinate goroutines, cancellation, bounded work, and result handling
  without turning the code into chaos

Suggested archetype:

- `parallel-fetch-pipeline`
- `download-orchestrator`

### Project 6: Architecture Phase Capstone

Placement:

- after Section 14

Role:

- prove testability, observability, packaging, deployment awareness, and service lifecycle thinking

Suggested archetype:

- `production-grade backend`
- `service-hardening lab`

### Project 7: Professional Practice Finalization

Placement:

- after Section 15

Role:

- apply generation and automation to a prior project so learners see tooling as leverage, not as a
  detached topic

Suggested archetype:

- generated interfaces, SQL access, or mock workflows added to the architecture capstone

## Reuse Strategy

v2 should prefer progressive reuse over total project reset.

Recommended pattern:

- mini-projects stand alone
- later phases may evolve or harden a previous project when that creates genuine continuity
- final capstone quality work can build on the architecture project instead of inventing a brand new
  product

## Project Rules

A v2 project should:

- validate a real curriculum jump
- have explicit scope and success criteria
- expose meaningful engineering boundaries
- avoid unrelated feature creep
- be supportable by the learner's actual prerequisites

## Deliverables Per Project

Every serious project should include:

- a short learner-facing README
- requirements
- verification instructions
- starter scaffolding when self-implementation is expected
- quality expectations appropriate to its phase

## Open Design Questions

These still need review:

- whether the architecture capstone should remain the single flagship project or become one of
  several equal capstones
- how much domain continuity we want between earlier projects and later service work
- whether every phase needs a mandatory project or whether some phases are better served by a strong
  checkpoint plus one larger follow-up project
