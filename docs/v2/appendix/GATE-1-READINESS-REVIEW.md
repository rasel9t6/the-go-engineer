# Gate 1 Readiness Review

Status: Ready for prototype work  
Date: 2026-04-07

## Purpose

This review records whether the v2 planning phase is coherent enough to begin the structural
prototype without inventing new system rules mid-implementation.

It does not approve broad curriculum migration.
It only answers whether the planning stack is strong enough to open prototype work under `#84`.

## Readiness Call

The current planning stack is strong enough to begin the prototype phase.

The key reason is that the repo now has a coherent first-pass system for:

- learner model
- content types
- lesson contract
- exercise and project roles
- learning paths
- docs navigation
- schema direction
- QA and validation expectations

The planning audit also confirmed that the remaining top-level spec gaps have now been filled:

- `13-CONTRIBUTOR-WORKFLOW.md`
- `16-RELEASE-ROLLOUT.md`

The few remaining future-facing questions are now covered by explicit working decisions or are
deferred beyond the first prototype on purpose.

## Evidence Reviewed

The readiness call is based on the current public planning set:

- `docs/v2/BIBLE.md`
- `docs/v2/00-VISION.md`
- `docs/v2/01-CURRICULUM-PHILOSOPHY.md`
- `docs/v2/02-LEARNER-MODEL.md`
- `docs/v2/03-CONTENT-TYPE-SYSTEM.md`
- `docs/v2/04-LESSON-SPEC.md`
- `docs/v2/05-EXERCISE-BANK-SYSTEM.md`
- `docs/v2/06-PROJECT-LADDER.md`
- `docs/v2/07-LEARNING-PATHS.md`
- `docs/v2/08-CURRICULUM-MAP.md`
- `docs/v2/09-FOLDER-STRUCTURE.md`
- `docs/v2/10-DOCS-NAVIGATION-SYSTEM.md`
- `docs/v2/11-CURRICULUM-SCHEMA.md`
- `docs/v2/12-MIGRATION-MAP-V1-TO-V2.md`
- `docs/v2/13-CONTRIBUTOR-WORKFLOW.md`
- `docs/v2/14-QA-VALIDATION.md`
- `docs/v2/15-IMPLEMENTATION-ROADMAP.md`
- `docs/v2/16-RELEASE-ROLLOUT.md`
- GitHub issues `#79` through `#88`

## What Is Ready

### 1. The System Direction Is Coherent

v2 is no longer just a bundle of ideas.
It now has a stable planning thesis:

- preserve the strong v1 topic order
- redesign the training system around that content
- migrate incrementally instead of doing a hidden rewrite

### 2. The Main Content Contracts Exist

The planning set now defines:

- what content types exist
- what a lesson is
- how exercises, checkpoints, mini-projects, and capstones differ
- how learning paths reuse the same curriculum objects

### 3. The Docs Model Is Coherent

The planning set now has a clear answer for:

- what the root README should do
- what the learning-path guide should do
- what the curriculum map should do
- why section READMEs are required

### 4. The Metadata Direction Is Strong Enough

The schema draft is mature enough to support a prototype slice.
The staged recommendation is also clear:

- keep `curriculum.json` active during early migration
- prototype richer v2 metadata in parallel

### 5. The QA Model Is Strong Enough

The repo now has a realistic split between:

- validator truth
- CI truth
- human review truth

That is enough to prototype without pretending every quality concern can be automated.

### 6. The GitHub Planning Structure Exists

The project board, milestones, and planning/prototype issue hierarchy are already in place.

## Prototype-Safe Working Decisions

These decisions are strong enough to guide the prototype now:

- keep the existing top-level section numbering and broad topic progression for the first migration
  waves
- use one curriculum with three canonical paths: Full Path, Bridge Path, and Targeted Path
- require a section README as a first-class learner navigation surface
- keep `curriculum.json` active while designing a richer v2 schema in parallel
- keep the validator focused on structural and metadata truth before expanding into more ambitious
  checks

## Remaining Deferred Questions

There are still future questions beyond Gate 1, but they are no longer prototype blockers.

Examples:

- whether coverage should become a harder release gate later
- whether the validator should split after prototype and early alpha work
- whether some doc placement changes make sense after real migration waves exist

Those are now rollout-stage refinements, not reasons to block the prototype.

## Prototype Constraints

Opening the prototype phase does not authorize broad migration.

Prototype work should follow these constraints:

- do not rewrite multiple live sections
- do not mass-rename top-level directories
- do not redesign the whole schema during the first prototype slice
- do not treat the prototype as public curriculum promotion

The prototype should prove the system, not expand scope.

## Prototype Success Criteria

The prototype should demonstrate one coherent slice that includes:

- one canonical section outline
- one canonical lesson
- one canonical exercise
- one canonical checkpoint
- one canonical mini-project
- one canonical metadata example
- one section-level navigation surface
- QA rules that can meaningfully check the slice

If that slice works, the planning stack is validated in practice rather than only on paper.

## Branching Recommendation For Prototype Work

Because the prototype is still planning-phase work, it should not land directly on learner-facing
`main`.

Recommended approach:

- keep prototype planning and non-learner-facing work on `planning/v2`
- if the prototype needs isolated experimentation, use a short-lived branch that merges back into
  `planning/v2`
- only promote prototype results to `main` after the system is validated and translated into
  learner-facing implementation work

## Recommended Next Move

Open the prototype phase under `#84`.

The immediate next issue should be:

1. `#85` choose and define the canonical prototype section outline

Recommended prototype target:

- start with a mid-curriculum section that is rich enough to prove lesson, exercise, checkpoint,
  and mini-project behavior without requiring full application-scale infrastructure

Section 04 is a strong candidate because it is conceptually important, structurally rich, and still
small enough to prototype cleanly.

## Bottom Line

The planning phase is not "finished forever," but it is now complete enough for the prototype phase
to begin without hidden system gaps.

The honest move now is:

- treat Gate 1 as passed
- begin the structural prototype
- use the prototype to validate the system in practice before broad migration
