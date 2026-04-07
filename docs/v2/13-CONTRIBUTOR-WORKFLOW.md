# V2 Contributor Workflow

## Purpose

v2 needs a contributor workflow that matches the repo boundary model and the staged migration plan.

This document defines how contributors and maintainers should move work from planning to prototype
to learner-facing implementation without polluting `main` or bypassing the design system.

## Core Rule

Not all repo work belongs on the same branch.

Contributors should choose the branch and workflow based on the kind of work they are doing:

- planning work
- prototype work
- learner-facing implementation
- stable v1 maintenance

## Repository Boundary Model

Use this boundary model consistently.

### `main`

Purpose:

- learner-facing v2 development line

Use for:

- approved learner-facing curriculum changes
- public docs that learners need now
- public tooling and validation scripts required to keep the repo healthy

Do not use for:

- rough planning docs
- internal decision logs that are not ready for public learner-facing relevance
- structural prototypes that are still proving the system

### `release/v1`

Purpose:

- stable v1 maintenance line

Use for:

- v1 bug fixes
- urgent stable-doc corrections
- current-user support changes

### `planning/v2`

Purpose:

- public v2 planning and prototype line

Use for:

- Bible docs
- migration strategy
- planning specs
- structural prototype work that is not yet learner-facing implementation

This is the correct branch for most current v2 planning work.

### Private Maintainer Repo

Purpose:

- internal-only maintainer operations

Use for:

- private notes
- maintainer automation
- review and release rehearsal notes
- confidential or non-public workflow material

Do not use it as the source branch for public PRs.

## Contribution Types And Where They Go

| Change Type | Base Branch | Example Branch |
| ----------- | ----------- | -------------- |
| v2 planning doc | `planning/v2` | `docs/v2-path-rules` |
| v2 structural prototype | `planning/v2` | `prototype/section-04-outline` |
| learner-facing v2 implementation | `main` | `feat/section-04-checkpoint` |
| v2 bug fix | `main` | `fix/run-command-doc` |
| v1 stable fix | `release/v1` | `fix/v1-curriculum-path` |
| maintainer-private workflow | private repo | local/private branch |

## Issue-First Rule

For v2 work, contributors should start from an issue whenever the change affects:

- curriculum structure
- planning docs
- prototype scope
- section migration
- release rollout

Planning and prototype work should not grow from untracked chat memory alone.

## Workflow By Phase

### Phase A: Planning Work

Use when:

- defining the system
- refining the Bible
- resolving migration strategy
- defining contracts for lessons, exercises, docs, schema, or QA

Workflow:

1. work from `planning/v2`
2. update the relevant planning docs
3. link the work to the planning issue
4. keep learner-facing `main` untouched

### Phase B: Prototype Work

Use when:

- proving the system with one canonical slice
- testing section README rules
- testing lesson/exercise/checkpoint/project contracts
- testing metadata and validator expectations in practice

Workflow:

1. branch from `planning/v2`
2. keep the prototype tightly scoped
3. prove the system without broad migration
4. merge prototype work back into `planning/v2`
5. only translate proven outcomes into `main` later

### Phase C: Learner-Facing Implementation

Use when:

- a planning or prototype rule has been validated
- the change belongs in the public learner product now

Workflow:

1. branch from `main`
2. implement only the approved learner-facing scope
3. run the public validation checks
4. open a PR back to `main`

### Phase D: Stable v1 Fixes

Use when:

- current stable users need a bug fix or correction

Workflow:

1. branch from `release/v1`
2. fix the stable issue
3. merge back to `release/v1`
4. cherry-pick forward to `main` if the same fix belongs there

## Pull Request Rules

### Planning And Prototype PRs

Planning and prototype PRs should target:

- `planning/v2`

They should include:

- issue reference
- what changed in the planning system
- what decisions were clarified
- what remains intentionally deferred

### Learner-Facing PRs

Learner-facing implementation PRs should target:

- `main` by default
- `release/v1` only for v1-only fixes

They should not include:

- stray planning-only docs
- internal notes
- prototype scaffolding that is not ready for learners

## Merge Rules

Maintainers should continue using:

- Squash and Merge for public PRs

For fixes that belong in more than one supported line:

1. merge into the branch that needs the fix first
2. use `git cherry-pick -x` to propagate it

Do not use branch-sync merges between supported lines.

## Validation Expectations By Work Type

### Planning Doc Changes

Minimum expectations:

- links are correct
- planning docs remain coherent
- issue status reflects the work

### Prototype Changes

Minimum expectations:

- prototype scope stays narrow
- related planning docs stay aligned
- metadata, docs, and sample structure agree

### Learner-Facing Changes

Minimum expectations:

- CI checks pass
- validator passes
- docs and commands are real
- the change matches the approved migration wave

## Reviewer Responsibilities

Reviewers should check different things depending on the work type.

### For Planning PRs

Check:

- coherence with the Bible
- whether a real decision was made
- whether open questions are captured instead of hidden

### For Prototype PRs

Check:

- whether the prototype proves the intended rule
- whether the slice is too large
- whether the work is still planning-phase rather than accidental migration

### For Learner-Facing PRs

Check:

- learner clarity
- runnable correctness
- path and navigation integrity
- fit with the migration wave

## Definition Of Done

The contributor workflow is ready when:

- contributors can tell which branch to use for which kind of work
- planning work no longer leaks into learner-facing `main`
- prototype work has a clear home before implementation
- v1 fixes and v2 work do not collide unnecessarily
- maintainers can explain the correct PR target without improvising

## Working Recommendation

For the current stage of v2:

- keep planning and prototype work on `planning/v2`
- keep learner-facing implementation on `main`
- keep stable fixes on `release/v1`
- keep private maintainer operations in the private maintainer repo
- do not skip from planning directly into broad curriculum migration
