# Rubric And Checkpoint Template

## Purpose

This template defines the default shape for beta checkpoint work, rubric-heavy review tasks, and
flagship checkpoints.

Use it when a task must prove readiness, judgment, or integrated engineering quality instead of
only asking the learner to finish a guided exercise.

## When To Use This Template

Use this template for:

- expert-layer review tasks
- expert-layer diagnosis tasks
- redesign-under-constraints tasks
- flagship project checkpoints
- rubric-verified stage milestones

Do not use this template for:

- simple drills
- first-pass guided exercises
- items where scaffolding is still the main teaching need

## Core Rule

Every rubric-heavy item must make four things explicit:

1. what the learner is being asked to do
2. what evidence counts as proof
3. what quality dimensions matter
4. what "done" means

## Canonical Structure

Use this as the default shape:

~~~md
# Item Title

## Mission

One short paragraph explaining the engineering goal.

## Type

- review task / diagnosis task / redesign task / flagship checkpoint

## Level

- foundation / core / stretch / production

## Prerequisites

- prior stage or milestone surfaces the learner should already understand

## Task

1. the concrete thing the learner must inspect, explain, or change
2. the constraint or scope boundary that keeps the task honest
3. the expected proof surface: note, explanation, code change, review comment set, or checkpoint outcome

## Evidence

- what output or explanation the learner must provide
- what code, behavior, or system surface the learner must point to
- what should be observable instead of hand-waved

## Rubric

### 1. Correctness

What must be true or working.

### 2. Completeness

What required pieces must be covered.

### 3. Boundary Handling

What seams, constraints, or non-happy-path behavior must be handled carefully.

### 4. Code Quality

What readability, structure, or design expectations matter.

### 5. Verification Discipline

What proof, checks, or evidence make the work believable.

### 6. Explanation Quality

What trade-off or reasoning quality is expected when the task asks for diagnosis or review judgment.

## Common Weak Answers

- one realistic shallow answer pattern
- one sign that the learner is naming problems without evidence
- one sign that the learner is proposing changes without trade-off reasoning

## Next Step

Where the learner should go next after this checkpoint or pressure task.
~~~

## Rubric Use Rules

The rubric should not be turned into fake precision.

Use it to answer:

- what absolutely must work
- what quality expectations matter
- what evidence counts
- what weak work usually looks like

## Flagship Checkpoint Notes

For flagship work, checkpoint prompts should usually include:

- system area in scope
- change budget or constraint
- expected proof of behavior
- expected explanation of trade-offs

## Expert Layer Notes

For expert-layer work, the best prompts usually force one of these:

- ranking risks by importance
- diagnosing a failure from evidence
- choosing a redesign under constraints
- defending one improvement over another

## Success Signal

This template is working when maintainers can create a new checkpoint or pressure task without
guessing:

- what the learner is meant to submit
- what the rubric should judge
- what proof is required
- how to distinguish shallow answers from real engineering judgment
