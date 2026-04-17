# The Go Engineer Curriculum Blueprint

## Purpose

This document defines how the 21-section v2 architecture should behave as a learning system.

It is not just a list of topics.
It is the curriculum contract for how we teach:

- from zero programming knowledge
- through Go fluency
- into software-engineering judgment

## Core Promise

The Go Engineer should help a learner move from:

- "I can copy code"

to:

- "I understand what this line does"
- "I can explain why this design exists"
- "I can predict what breaks"
- "I can build and operate a real system"

## Non-Negotiable Teaching Rules

### 1. Doc first, then code

Every learner-facing lesson teaches through `README.md` first.
The learner should understand the mission before opening `main.go`.

### 2. Code is never skipped

We do not replace code with prose.
We explain the code, then run the code, then modify the code.

### 3. Zero magic

Each section teaches only what has been earned.
If a concept depends on later ideas, it belongs later or must be reduced to a clearly labeled preview.

### 4. Explanation should answer how, why, and what changes

Good teaching surfaces explain:

- what this line or block does
- why it exists
- what would change if we changed it
- what mistake a learner is likely to make here

### 5. Engineering depth must be stage-aware

We do want:

- design thinking
- failure thinking
- production relevance
- debugging instincts

But we add them at the right layer.
We do not dump senior-level pressure framing into beginner lessons just to sound impressive.

## Phase-Level Blueprint

The curriculum is split into 5 phases. See `ARCHITECTURE.md` for the exact 21 sections within these phases.

### Phase 1 & 2: Syntax & Engineering Foundation

These sections must feel safe, explicit, and zero-magic.

Required elements:

- mission
- mental model
- literal or near-literal walkthroughs
- clean runnable code

Avoid:

- premature scale pressure
- advanced security catalogs
- abstract design jargon before the learner has concrete examples

### Phase 3 & 4: Production & Systems Engineering

These sections should start increasing engineering judgment.

Add more of:

- trade-off explanations
- failure cases
- safer defaults
- tests and verification surfaces
- performance and maintainability reasoning
- architecture trade-offs
- production notes

### Phase 5: Flagship Project (GoScale)

This phase should carry heavier:

- integrated project proof
- production deployment
- operational pressure

## Canonical Lesson Contract

For learner-facing lessons, the default shape is:

```text
lesson/
├── README.md
├── main.go
├── _bad/
│   └── main.go            (optional)
├── _starter/
│   └── main.go            (optional)
└── main_test.go           (optional)
```

### Required README sections

Each lesson README should include, as appropriate:

- mission
- prerequisites
- mental model
- run instructions
- code walkthrough
- `Try It`
- next step

### Required source-file behavior

Source files should stay readable and runnable.
They should not become giant essays.

Use inline comments for:

- non-obvious behavior
- mutation or boundary traps
- subtle runtime implications

Do not use code headers as the main teaching surface.

## Canonical Milestone Contract

Every section needs proof, not just lessons.

A milestone should usually provide:

- clear README instructions
- a runnable completed solution
- a starter surface when the learner is expected to implement
- tests when the behavior should be provable

## How To Add New Lessons Without Breaking The Architecture

If the curriculum needs more depth:

1. Add the lesson inside an existing section.
2. Keep the learner-facing section count at exactly 21.
3. Update `ARCHITECTURE.md` if the scope of the section expands.

Do not solve content growth by inventing a new root-level section unless the entire architecture is being reworked intentionally.

## Bottom Line

The Go Engineer should feel like one coherent engineering learning system.

The 21 sections give us the public spine.
The README-first teaching contract gives us the delivery standard.
Future expansion should deepen the sections we have, not fragment the learner path again.
