# V2 Validator Expansion Plan

## Purpose

This document turns the v2 QA model into an implementation-ready plan for the first validator
expansion.

It is the working answer for issue `#95`.
Its job is to tell maintainers:

- which rules belong in the first v2 validator pass
- which rules are already proven by the Section 04 prototype
- which checks should stay deferred
- how to extend `scripts/validate_curriculum.go` without inventing a second validator system

## Working Boundary

The first validator expansion should enforce structural truth, not teaching quality.

It should prove:

- the metadata is coherent enough to trust
- the declared item paths exist
- the declared command targets resolve
- the item graph is connected correctly
- the item type and verification contracts are not self-contradictory

It should not try to prove:

- whether a lesson explanation is good
- whether an exercise is interesting enough
- whether a checkpoint is pedagogically difficult enough
- whether a project scope is perfectly calibrated

Those remain review concerns.

## Current Baseline

Today `scripts/validate_curriculum.go` already does two useful things:

1. validates that `curriculum.json` lesson paths exist
2. validates that `go run` and `go test` path references in Go files and numbered section READMEs
   point to real targets

That behavior should stay intact.
The first v2 work should extend the current tool, not replace it.

## First Expansion Goal

The first v2 validator pass should be able to validate the first migrated v2 section without
guessing.

The Section 04 prototype proves the intended contract across four item shapes:

- `FEP.4` lesson
- `FEP.E1` exercise
- `FEP.C1` checkpoint
- `FEP.P1` mini-project

If the validator cannot check those shapes meaningfully, the first migration wave is still too
fragile.

## Activation Rule

The first v2 validator pass should stay additive while the repo is still in planning mode.

Implementation rule:

- keep the current v1 checks running by default
- add a v2-aware pass that activates only when the chosen v2 metadata input exists
- do not fail planning-only branches that have no migrated v2 metadata yet

This keeps the validator honest during the migration window.

Issue `#96` will decide the exact metadata source shape.
Issue `#95` defines the rules that source must support.

## Priority Rule Set

The first implementation should use the following priorities.

### Priority 0: Preserve Existing Truth

These are mandatory carry-forward checks:

- `curriculum.json` lesson paths still validate
- `go run` and `go test` references in live repo surfaces still resolve

The first v2 validator PR should not regress existing v1 coverage.

### Priority 1: Add High-Value Metadata And Graph Checks

These are the first v2-specific rules to add.

| Rule | What The Validator Must Prove | Prototype Proof |
| ---- | ----------------------------- | --------------- |
| unique ids | section ids and content-item ids are unique | `s04-prototype`, `FEP.4`, `FEP.E1`, `FEP.C1`, `FEP.P1` are distinct |
| section linkage | every item's `section_id` resolves to a real section | all prototype items resolve to `s04-prototype` |
| section path alignment | each item path stays under its section's declared `path_prefix` | each Section 04 prototype item lives under `04-functions-and-errors/...` |
| prerequisite linkage | each prerequisite id resolves to a real item or allowed prior section id | `FEP.4` depends on `FEP.3`; milestone items depend on earlier `FEP.*` items |
| next-item linkage | each `next_items` entry resolves to a real item id or allowed section id | `FEP.4 -> FEP.E1 -> FEP.5 -> FEP.C1 -> FEP.P1 -> s09` |
| enum validity | `type`, `subtype`, `level`, and `verification_mode` use allowed values | `lesson/pattern/core/run`, `exercise/core/mixed`, `mini_project/core/mixed` |

### Priority 2: Add Path And Command Truth

These rules should land in the same first expansion if possible because they are high-signal and
already match the validator's strengths.

| Rule | What The Validator Must Prove | Prototype Proof |
| ---- | ----------------------------- | --------------- |
| item path exists | each item's declared `path` exists relative to repo root | each Section 04 prototype item declares a concrete section-local path |
| run command resolves | every `run_command` points to a real package or file target | `FEP.4`, `FEP.E1`, `FEP.C1`, and `FEP.P1` all declare concrete `go run` targets |
| test command resolves | every non-empty `test_command` points to a real package target | supported even though the prototype leaves `test_command` empty |
| starter path resolves | every non-empty `starter_path` exists | `FEP.E1` declares `_starter/`; the others do not |

### Priority 3: Add Verification And Type Contracts

These rules should still be part of the first implementation-ready scope, but they should only
check what the repo can prove cheaply.

| Rule | What The Validator Must Prove | Prototype Proof |
| ---- | ----------------------------- | --------------- |
| `run` contract | `verification_mode: run` items declare a non-empty `run_command` | `FEP.4` |
| `test` contract | `verification_mode: test` items declare a non-empty `test_command` | future-facing rule |
| `rubric` contract | `verification_mode: rubric` items include a learner-facing README | future-facing rule |
| `mixed` contract | `verification_mode: mixed` items include at least one runnable/testable command and a learner-facing README | `FEP.E1`, `FEP.C1`, `FEP.P1` |
| exercise README | exercises include a learner-facing `README.md` under the item path | `FEP.E1` |
| checkpoint README | checkpoints include a learner-facing `README.md` under the item path | `FEP.C1` |
| mini-project README | mini-projects include a learner-facing `README.md` under the item path | `FEP.P1` |
| mini-project entrypoint | a mini-project with a declared `run_command` exposes the declared entrypoint inside the project layout | `FEP.P1` uses `cmd/task-journal` |

## Required Metadata Minimums

The first v2 pass should fail fast when required structural fields are missing.

### Section Minimums

Require at least:

- `id`
- `number`
- `slug`
- `title`
- `path_prefix`

### Content Item Minimums

Require at least:

- `id`
- `section_id`
- `slug`
- `title`
- `type`
- `level`
- `verification_mode`
- `path`
- `prerequisites`
- `next_items`

The first pass may treat `subtype`, `run_command`, `test_command`, `starter_path`, and `tags` as
conditionally required depending on item type or verification mode.

## Prototype-Backed Acceptance Criteria

The first validator expansion should be considered ready when it can enforce checks equivalent to
the following:

### Lesson Acceptance Criteria

Given a lesson record shaped like `FEP.4`, the validator should fail when:

- `type` is not `lesson`
- `subtype` is not an allowed lesson subtype
- `run_command` is empty even though `verification_mode` is `run`
- the declared lesson path does not exist
- `next_items` points to a missing item

### Exercise Acceptance Criteria

Given an exercise record shaped like `FEP.E1`, the validator should fail when:

- the exercise path does not exist
- `starter_path` is declared but missing
- `verification_mode: mixed` has no runnable or testable command
- `README.md` is missing from the exercise path
- `next_items` points to a missing item

### Checkpoint Acceptance Criteria

Given a checkpoint record shaped like `FEP.C1`, the validator should fail when:

- the checkpoint path does not exist
- `verification_mode: mixed` has no runnable or testable command
- `README.md` is missing from the checkpoint path
- `next_items` points to a missing item

### Mini-Project Acceptance Criteria

Given a mini-project record shaped like `FEP.P1`, the validator should fail when:

- the project path does not exist
- `README.md` is missing from the project path
- the declared `run_command` target does not exist
- the project layout does not expose the declared entrypoint
- `next_items` points to a missing item or section

## Deferred Checks

The following checks are valuable, but they should stay out of the first validator expansion:

- judging whether objectives are well written
- judging whether `production_relevance` is persuasive enough
- checking exact README section wording beyond existence-based structural rules
- checking broader section README and root-navigation integrity beyond the first migrated slice
- validating every Markdown link across the full repo
- executing lesson or project commands as part of the validator
- checking whether checkpoint difficulty is appropriate
- adding drill-specific or capstone-specific rules before those shapes exist in live migrated form
- enforcing coverage or benchmark thresholds
- splitting the validator into multiple tools

These are good later candidates once the first migrated sections exist and the team has seen real
failure modes.

## Implementation Scope For `scripts/validate_curriculum.go`

The first validator expansion should keep one entrypoint:

- `go run ./scripts/validate_curriculum.go`

Recommended implementation shape:

1. keep the current v1 curriculum-path pass
2. keep the current run-path scanning pass
3. add a new v2 metadata pass behind an existence check for the chosen v2 metadata source
4. reuse one shared command-target resolver for both run-path scanning and v2 command validation
5. accumulate all validation errors before exiting so authors get one useful report

The first implementation should not:

- create a second standalone validator binary
- require full v2 migration before the tool can run
- overfit to planning-only docs instead of live migrated content

## Recommended Error Style

When the new v2 pass fails, the validator should report errors in a path- and id-aware way.

Examples:

- `Invalid next item: FEP.C1 -> FEP.P1`
- `Missing README for checkpoint: 04-functions-and-errors/8-reliable-batch-processor-checkpoint`
- `Invalid starter path: FEP.E1 -> 04-functions-and-errors/6-safe-operations-runner/_starter`
- `Invalid run command target: FEP.P1 -> ./04-functions-and-errors/9-task-journal-cli/cmd/task-journal`

This keeps failures actionable for maintainers.

## Success Signal

Issue `#95` should be considered successful when maintainers can say:

- we know exactly which rules belong in the first v2 validator PR
- we know which prototype surfaces prove those rules are worth enforcing
- we know which checks are intentionally deferred
- we can extend the existing validator without redesigning the tooling stack first
