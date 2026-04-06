# V2 Implementation Roadmap

Status: Draft execution roadmap.

## Governing Rule

No large-scale repo migration starts before the planning docs and the structural prototype are both
approved.

This roadmap is intentionally front-loaded with design work so we do not repeat v1 drift at v2
scale.

## Workstreams

v2 should be executed through six workstreams:

1. Bible and planning docs
2. Structural prototype
3. Tooling and templates
4. Docs and curriculum schema
5. Section migration waves
6. Release and rollout

## GitHub Planning Model

### Milestones

Use these milestones:

- `v2 planning`
- `v2 prototype`
- `v2 alpha`
- `v2 beta`
- `v2 rc`
- `v2.0.0`

### Labels

Use the labels already created for repo operations, then add v2 planning labels if needed:

- `v2`
- `backlog`
- `prototype`
- `curriculum`
- `exercise-system`
- `migration`
- `release-blocker`

### Project

Use one GitHub Project for the v2 program.

Recommended status columns:

- Backlog
- In Design
- Prototype Ready
- Ready For Breakdown
- In Progress
- In Review
- Approved
- Scheduled For Alpha
- Done

Recommended fields:

- Section
- Content Type
- Milestone
- Target Branch
- Migration Type
- Prototype Required
- Owner

## Issue Hierarchy

Create epics first, then child issues.
Do not create every lesson issue across all sections before the prototype is approved.

### Epic 0: V2 Bible And Strategy

- define v2 vision and non-goals
- define curriculum philosophy
- define exercise bank system
- define curriculum map
- define migration map
- define implementation roadmap

### Epic 1: V2 Prototype

- prototype one canonical section
- prototype one canonical lesson
- prototype one canonical guided exercise
- prototype one canonical mini-project
- prototype one curriculum metadata example
- review and approve prototype

### Epic 2: Tooling And Templates

- define lesson template
- define exercise template
- define project template
- extend validator scope
- prepare curriculum schema changes

### Epic 3: Docs And Navigation

- rewrite v2 learner navigation docs
- design section README standard
- design learning path standard
- draft migration guide shell

### Epic 4+: Section Migration Epics

Create one epic per active section wave.
Each section epic should include:

- section scope and outcomes
- lesson list review
- exercise and checkpoint plan
- mini-project linkage
- docs and navigation updates
- metadata updates
- QA and validation pass

### Final Epic: Release And Rollout

- alpha release plan
- beta freeze criteria
- RC checklist
- final v2.0.0 checklist
- v1 support policy and migration messaging

## Execution Phases

### Phase 0: Planning

Deliverables:

- first draft Bible docs
- open questions list
- issue hierarchy approved

Milestone: `v2 planning`

### Phase 1: Prototype

Deliverables:

- one full prototype section outline
- one prototype lesson
- one prototype exercise
- one prototype mini-project
- one prototype metadata example

Milestone: `v2 prototype`

### Phase 2: Infrastructure

Deliverables:

- v2 templates
- validator upgrades
- documentation skeleton
- migration-ready issue backlog

Milestone: `v2 alpha`

### Phase 3: Alpha Migration Waves

Recommended wave order:

1. Sections 01-04
2. Sections 05-08
3. Sections 09-12
4. Sections 13-15

Alpha releases should be cut from `main` only when a wave produces a meaningful public improvement.

### Phase 4: Beta Stabilization

At the point where planned v2 scope is mostly complete:

- cut `release/v2` from `main`
- stop feature expansion
- focus on consistency, migration docs, and bug fixes

Milestone: `v2 beta`

### Phase 5: Release Candidate

Focus on:

- migration guide completeness
- docs coherence
- validator coverage
- release notes and support messaging

Milestone: `v2 rc`

### Phase 6: Final Release

Publish:

- `v2.0.0`
- migration guide
- updated README and learner-facing docs
- clear v1 support policy

Milestone: `v2.0.0`

## Branch And Release Rules

During implementation:

- `main` is the active v2 line
- `release/v1` stays stable for current users
- `release/v2` is created only at feature freeze

For fixes that belong in both active lines:

1. merge to the branch that needs the fix first
2. squash merge the PR
3. `git cherry-pick -x` the merged commit to the other supported branch

## Immediate Next Actions

The next practical steps after this draft are:

1. review and refine the Bible docs
2. create the planning and prototype GitHub issues
3. open a GitHub Project for v2 tracking
4. draft the first prototype section and lesson
5. refuse broad content migration until the prototype is approved

## Definition Of Readiness For Implementation

v2 is ready for real migration only when:

- the Bible documents are good enough to guide contributors
- the prototype demonstrates the lesson, exercise, and mini-project system
- the issue hierarchy is created in GitHub
- milestones and project views are in place
- the first migration wave has an approved scope
