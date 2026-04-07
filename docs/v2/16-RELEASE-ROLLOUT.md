# V2 Release Rollout

## Purpose

v2 needs a release rollout plan that lets the project evolve publicly without disrupting stable v1
learners.

This document defines how planning becomes prototype work, how prototype work becomes public v2
prereleases, and how stable v2 should eventually replace v1 as the default learner line.

## Core Rule

v2 should be released incrementally and honestly.

That means:

- stable v1 users should always know where the safe line is
- v2 users should know when they are on alpha, beta, RC, or final
- prerelease work should not be marketed as fully stable

## Supported Release Lines

Use these long-lived branches consistently:

- `main`: active v2 development and alpha prerelease line
- `release/v1`: stable v1 support line
- `release/v2`: created later when v2 enters beta and feature freeze

## Rollout Stages

### Stage 0: Planning

Status:

- public planning happens on `planning/v2`

Outputs:

- Bible
- issue hierarchy
- prototype definition

No learner-facing migration happens here.

### Stage 1: Prototype

Status:

- structural prototype work continues on `planning/v2`

Outputs:

- one validated system slice
- confirmed lesson, exercise, checkpoint, project, metadata, and navigation behavior

This stage still does not publish learner-facing v2 releases.

### Stage 2: Alpha

Status:

- approved learner-facing migration work lands on `main`

Outputs:

- meaningful v2 improvements land section by section
- public alpha prereleases can be tagged from `main`

Tag pattern:

- `v2.0.0-alpha.N`

Alpha rules:

- mark releases as prereleases
- do not mark them as the latest stable release
- be explicit that stable users can stay on `release/v1`

### Stage 3: Beta

Status:

- create `release/v2` from `main` once planned v2 scope is mostly complete and feature expansion
  pauses

Outputs:

- stabilization work
- consistency fixes
- docs cleanup
- migration guide maturity

Tag pattern:

- `v2.0.0-beta.N`

### Stage 4: Release Candidate

Status:

- `release/v2` continues as the stabilization line

Outputs:

- final migration messaging
- near-final docs
- final release blockers only

Tag pattern:

- `v2.0.0-rc.N`

### Stage 5: Final

Status:

- release `v2.0.0` from `release/v2`

Outputs:

- final GitHub release
- migration guide
- updated learner-facing docs
- explicit v1 support policy

## Tagging Rules

Use these tag rules:

- alpha tags come from `main`
- beta, RC, and final tags come from `release/v2`
- v1 patch tags come from `release/v1`

Do not publish v2 prerelease tags from `release/v1`.

## Communication Rules

Each rollout stage should have clear public messaging.

### During Planning And Prototype

Tell users:

- v1 remains the stable line
- v2 planning is public
- no migration is required yet

### During Alpha

Tell users:

- v2 is usable for early adopters
- the curriculum is still evolving
- stable learners should stay on `release/v1` if they want the safest path

### During Beta And RC

Tell users:

- v2 structure is mostly settled
- migration guidance is available
- final stability work is underway

### During Final Release

Tell users:

- v2 is now the recommended stable line
- where to find the migration guide
- how long v1 will continue receiving support

## Migration Guide Timing

The migration guide should not wait until the final release.

Recommended timing:

- initial shell during alpha
- substantial draft during beta
- complete guide by RC
- final polished guide at `v2.0.0`

## V1 Support Policy

For the first v2 rollout:

- keep `release/v1` stable through planning, prototype, alpha, beta, and final release
- do not silently abandon v1 users
- announce a v1 support window once v2 final is public and the migration guide is complete

## Release Immutability Guidance

Published releases should be treated as permanent artifacts.

Recommended policy:

- use draft releases while preparing notes and assets
- publish immutable prereleases and finals once ready
- do not rewrite published stable release history casually

## Branch Promotion Rules

Use this promotion flow:

1. planning and prototype validation on `planning/v2`
2. approved learner-facing migration work on `main`
3. alpha tags from `main`
4. stabilization branch cut to `release/v2`
5. beta, RC, and final tags from `release/v2`

This keeps planning, active implementation, and stabilization separate.

## Release Readiness Gates

### Ready For Alpha

- prototype validated
- first migration wave scoped
- learner-facing docs honest about alpha status
- CI and validator remain healthy

### Ready For Beta

- planned v2 scope mostly complete
- migration guide substantially drafted
- docs/navigation model proven across active sections
- `release/v2` cut from `main`

### Ready For RC

- only release blockers remain
- docs, metadata, and repo paths are coherent
- known migration blockers are addressed

### Ready For Final

- migration guide is complete
- release notes are complete
- README and learning-path entry docs reflect v2 as the stable line
- v1 support messaging is explicit

## Definition Of Done

The rollout plan is ready when:

- every stage has a clear branch and tag strategy
- stable users always know the safe branch
- prerelease users know what line they are adopting
- migration guide timing is explicit
- the project can explain how v1 transitions to v2 without improvisation

## Working Recommendation

For the current phase:

- do not publish learner-facing v2 releases yet
- validate the system through prototype work first
- keep `release/v1` as the stable learner line
- move to alpha only after prototype results are translated into real learner-facing changes on
  `main`
