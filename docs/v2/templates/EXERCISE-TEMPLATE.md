# V2 Exercise Template

## Purpose

This document defines the canonical reusable guided-exercise template for v2.

It turns the approved exercise planning and prototype work into a maintainable authoring surface
that contributors can copy without improvising:

- exercise metadata
- README shape
- `_starter/` rules
- success criteria
- verification expectations
- validator implications

This template is derived primarily from:

- `docs/v2/05-EXERCISE-BANK-SYSTEM.md`
- `docs/v2/14-QA-VALIDATION.md`
- `docs/v2/prototype/EXERCISE-FEP-E1-SAFE-OPERATIONS-RUNNER.md`

## When To Use This Template

Use this template when the content item is:

- `type: exercise`

Use an exercise when:

- the learner must combine two to four prior lessons
- the learner benefits from constraints and scaffolding
- the goal is supported synthesis rather than final readiness proof

Do not use this template for:

- drills
- checkpoints
- mini-projects

## Canonical Exercise Directory Shape

Default guided-exercise layout:

```text
N-exercise-name/
  README.md
  main.go
  _starter/
    main.go
  main_test.go     # optional
  testdata/        # optional
```

The learner-facing README is required by default.

## Metadata Stub

Every v2 guided exercise should start with a metadata draft like this:

```json
{
  "id": "SX.E1",
  "section_id": "sNN",
  "slug": "exercise-slug",
  "title": "Exercise Title",
  "type": "exercise",
  "level": "core",
  "verification_mode": "mixed",
  "estimated_time": 60,
  "summary": "One-sentence description of the exercise outcome.",
  "objectives": [
    "Primary synthesis objective",
    "Optional supporting objective"
  ],
  "prerequisites": ["SX.3", "SX.4"],
  "production_relevance": "One concrete sentence about where this practice shape appears in real Go work.",
  "path": "NN-section-name/N-exercise-name",
  "run_command": "go run ./NN-section-name/N-exercise-name",
  "test_command": "",
  "starter_path": "NN-section-name/N-exercise-name/_starter",
  "next_items": ["SX.5"],
  "tags": ["exercise", "topic-a", "topic-b"]
}
```

## Metadata Field Notes

- `type` must always be `exercise`
- `verification_mode` should usually be `mixed`
- `starter_path` should be present when the exercise is truly scaffolded
- `next_items` should usually point to the next lesson, checkpoint, or mini-project linkage

## Canonical README Shape

Every v2 guided exercise README should include these sections:

1. mission
2. prerequisites
3. requirements
4. starter guidance
5. verification steps
6. success criteria
7. common failure modes
8. next step

That shape keeps the exercise practical and learner-facing without turning it into a second lesson.

## Canonical README Skeleton

Use this as the default exercise README shape:

~~~md
# Exercise Title

## Mission

What the learner is building and why it matters here.

## Prerequisites

- lesson or checkpoint ids the learner should already know

## Requirements

1. concrete behavior requirement
2. concrete behavior requirement
3. constraint or scope rule

## Starter Guidance

What `_starter/` gives the learner and what it intentionally does not solve.

## Verification

~~~bash
go run ./NN-section-name/N-exercise-name
~~~

Optional:

~~~bash
go test ./NN-section-name/N-exercise-name
~~~

## Success Criteria

- correct behavior for valid input
- expected failure handling is visible
- output or test results match the intended shape

## Common Failure Modes

- one realistic mistake learners often make
- one debugging hint that pushes the learner to inspect behavior honestly

## Next Step

Point to the next lesson, checkpoint, or mini-project.
~~~

## `_starter/` Rules

Guided exercises should usually include:

- `_starter/`

Use `_starter/` when:

- the exercise is the learner's first synthesis step after several lessons
- the structural skeleton would otherwise consume the learner's effort
- you want the learner to focus on decisions inside the implementation, not on bootstrapping

Do not use `_starter/` when:

- the work is supposed to validate independent readiness
- the item is actually a checkpoint or project
- the starter would solve the most important design decisions already

## What The Starter Should Provide

The starter may provide:

- core data shapes
- sample input setup
- TODO markers
- file and function skeletons
- high-level output expectations

The starter should not provide:

- the main algorithm
- completed error or verification flow
- the most important design boundary the learner is meant to create

## Success Criteria Rules

Exercise success criteria should:

- focus on observable behavior
- include at least one valid-flow expectation
- include at least one expected-failure expectation
- confirm the learner used the intended concept combination
- remain smaller than checkpoint pass criteria

Good exercise criteria support the learner without hiding the bar.

## Failure-Based Learning Rules

Guided exercises should include at least one meaningful failure surface.

That can be:

- invalid input
- unknown operation
- boundary misuse
- a case where the learner must classify or wrap an error

The goal is not to create a trick.
The goal is to stop the exercise from becoming pure happy-path copying.

## Scope Control Rules

Keep the exercise bounded:

- one local domain
- no unnecessary infrastructure
- no premature multi-package architecture unless it teaches something essential
- no mini-project-sized deliverable

If the exercise needs too much setup, either add better scaffolding or reduce the scope.

## Validator Notes

The validator should eventually enforce these exercise-specific checks:

- `type` is `exercise`
- `path` exists
- `starter_path` exists when declared
- `run_command` or `test_command` points to a real target
- `next_items` resolves
- a learner-facing README exists

The validator should not try to judge:

- whether the prompt is inspiring enough
- whether the constraints are pedagogically ideal
- whether the exercise is too easy for a specific audience

Those remain reviewer judgments.

## Success Signal

This template is working when a maintainer can copy it and produce a new v2 exercise without
guessing:

- what scaffolding belongs in `_starter/`
- what belongs in the README
- what counts as success criteria
- how much failure pressure the exercise should include
