# V2 Checkpoint Template

## Purpose

This document defines the canonical reusable checkpoint template for v2.

It turns the approved checkpoint planning and prototype work into a maintainable authoring surface
that contributors can copy without improvising:

- checkpoint metadata
- readiness criteria
- rubric/pass-criteria shape
- README structure
- validator implications

This template is derived primarily from:

- `docs/v2/05-EXERCISE-BANK-SYSTEM.md`
- `docs/v2/14-QA-VALIDATION.md`
- `docs/v2/prototype/CHECKPOINT-FEP-C1-RELIABLE-BATCH-PROCESSOR.md`

## When To Use This Template

Use this template when the content item is:

- `type: checkpoint`

Use a checkpoint when:

- the learner needs a real readiness signal before larger work
- the section has introduced enough material that supported practice is no longer enough
- the repo needs a visible progression gate

Do not use this template for:

- guided exercises
- mini-projects
- capstones

## Canonical Checkpoint Directory Shape

Default checkpoint layout:

```text
N-checkpoint-name/
  README.md
  main.go
  main_test.go     # optional
  testdata/        # optional
```

The README is required because the checkpoint must make "ready" mean something concrete.

## Metadata Stub

Every v2 checkpoint should start with a metadata draft like this:

```json
{
  "id": "SX.C1",
  "section_id": "sNN",
  "slug": "checkpoint-slug",
  "title": "Checkpoint Title",
  "type": "checkpoint",
  "level": "core",
  "verification_mode": "mixed",
  "estimated_time": 75,
  "summary": "One-sentence description of the readiness artifact.",
  "objectives": [
    "Primary readiness objective",
    "Optional supporting objective"
  ],
  "prerequisites": ["SX.4", "SX.5"],
  "production_relevance": "One concrete sentence about why this readiness shape matters in real Go work.",
  "path": "NN-section-name/N-checkpoint-name",
  "run_command": "go run ./NN-section-name/N-checkpoint-name",
  "test_command": "",
  "starter_path": "",
  "next_items": ["SX.P1"],
  "tags": ["checkpoint", "topic-a", "topic-b"]
}
```

## Metadata Field Notes

- `type` must always be `checkpoint`
- `verification_mode` should usually be `mixed` or `rubric`
- `starter_path` should normally be empty
- `next_items` should point to the milestone or section step the learner has now earned

## Canonical README Shape

Every v2 checkpoint README should include these sections:

1. checkpoint mission
2. prerequisites
3. required behavior
4. pass criteria
5. verification steps
6. common failure modes
7. next step

This keeps the checkpoint honest and prevents readiness from becoming vague.

## Canonical README Skeleton

Use this as the default checkpoint README shape:

~~~md
# Checkpoint Title

## Mission

What readiness this checkpoint is validating.

## Prerequisites

- lesson and exercise ids the learner should already understand

## Required Behavior

1. concrete behavior requirement
2. concrete behavior requirement
3. required boundary or design expectation

## Pass Criteria

- what must be correct
- what failure handling must be visible
- what scope rule keeps the checkpoint honest

## Verification

~~~bash
go run ./NN-section-name/N-checkpoint-name
~~~

Optional:

~~~bash
go test ./NN-section-name/N-checkpoint-name
~~~

## Common Failure Modes

- one common design mistake
- one common verification mistake
- one debugging hint that helps the learner inspect failure honestly

## Next Step

Point to the next mini-project, section, or milestone.
~~~

## `_starter/` Rules

Checkpoints should normally not include:

- `_starter/`

This is the default because a checkpoint is a readiness boundary, not a supported practice surface.

Only consider scaffolding when:

- setup burden would otherwise dominate the checkpoint
- the checkpoint is validating a later skill, not bootstrapping the environment

That should be rare.

## Pass-Criteria Rules

Checkpoint pass criteria should:

- be concrete and observable
- validate the section's real outcomes, not trivia
- include at least one expected-failure or recovery boundary when relevant
- distinguish correct behavior from merely runnable behavior
- stop short of mini-project sprawl

If the learner could pass by only producing a happy-path demo, the checkpoint is probably too weak.

## Failure Pressure Rules

Checkpoints should include meaningful pressure.

That can be:

- an expected failure class the learner must classify correctly
- a cleanup or boundary condition
- a recovered failure path
- a summary that must remain trustworthy even when some work fails

This is where the useful part of `failure-based learning` belongs in the v2 system:

- not as chaos
- not as trick questions
- but as proof that the learner can handle non-happy-path behavior intentionally

## Scope Control Rules

Keep the checkpoint bounded:

- one small domain
- one readable artifact
- no extra infrastructure unless the section genuinely taught it
- no milestone-sized deliverable

If it starts looking like a named product, it is probably a mini-project instead.

## Validator Notes

The validator should eventually enforce these checkpoint-specific checks:

- `type` is `checkpoint`
- `path` exists
- `run_command` or `test_command` points to a real target
- `next_items` resolves
- a learner-facing README exists
- visible pass criteria exist in the README

The validator should not try to judge:

- whether the checkpoint is stressful enough
- whether the pass criteria are optimally calibrated
- whether the section should have used a checkpoint at all

Those remain reviewer judgments.

## Success Signal

This template is working when a maintainer can copy it and produce a new v2 checkpoint without
guessing:

- whether `_starter/` belongs there
- what should count as pass criteria
- how much failure pressure is appropriate
- how to keep the checkpoint smaller than a mini-project
