# SL.5 Redaction Boundary Review

## Mission

Review the `SL.5` redaction milestone and explain whether sensitive-data handling is enforced in
the right place.

## Type

- review task

## Level

- stretch

## Prerequisites

- [8 Production Engineering](../../08-production-engineering.md)
- `SL.5`
- [rubric and checkpoint template](../../../templates/rubric-checkpoint-template.md)

## Task

1. Review the `SL.5` redaction surface and identify the two most important risks around log
   redaction boundaries.
2. Explain which risk matters most in a real production setting and why.
3. Propose one bounded improvement that would make the logging surface safer without turning it
   into overbuilt infrastructure.

## Evidence

- point to the specific logging boundary you are evaluating
- explain why the risk is real in terms of sensitive data, handler placement, or operational usage
- justify why the proposed improvement is the right next move

## Rubric

### 1. Correctness

The review should identify plausible operational risks in the logging/redaction flow.

### 2. Completeness

The learner should cover two risks, prioritize them, and propose one bounded improvement.

### 3. Boundary Handling

The review should discuss where redaction belongs and why that boundary matters.

### 4. Code Quality

The proposed change should improve safety or clarity without uncontrolled complexity.

### 5. Verification Discipline

The critique should rely on evidence from the current milestone surface.

### 6. Explanation Quality

The learner should explain why the top risk matters most under production pressure.

## Common Weak Answers

- treating redaction like a formatting preference instead of a boundary decision
- proposing a total rewrite without explaining the current risk clearly
- naming sensitive-data concerns without pointing to the actual surface

## Next Step

Continue to the redesign task at
[GS.3 shutdown-order redesign](./redesign-gs3-shutdown-order.md)
or move into the
[Flagship Project stage](../../10-flagship-project.md)
if you want system-scale pressure next.
