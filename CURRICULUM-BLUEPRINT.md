# The Go Engineer Curriculum Blueprint

This document defines how the locked v2.1 architecture should behave as a learning system.

If this file and [ARCHITECTURE.md](./ARCHITECTURE.md) disagree on public structure, `ARCHITECTURE.md` wins.

## Core Promise

The Go Engineer should help a learner move from:

- “I can copy code”

to:

- “I understand what this line does”
- “I can explain why this design exists”
- “I can predict what breaks”
- “I can build and operate a real system”

## Architecture Lock

The public curriculum spine is locked at 12 sections, `s00` through `s11`.

Do not solve content growth by inventing a new root-level section. Add depth inside existing sections unless the maintainer explicitly approves reopening architecture.

## Non-Negotiable Teaching Rules

### 1. README first, code second

Every learner-facing lesson teaches through `README.md` first.

The learner should understand the mission before opening `main.go`.

### 2. Code is never skipped

Do not replace code with prose. Explain the code, run the code, then modify the code.

### 3. Zero magic

Each section teaches only what has been earned.

If a concept depends on later ideas, either move it later or add a clearly labeled forward reference.

### 4. Explanation must answer how, why, and what changes

Good teaching surfaces explain:

- what this line or block does
- why it exists
- what would change if we changed it
- what mistake a learner is likely to make here

### 5. Engineering depth must be stage-aware

Add design thinking, failure thinking, production relevance, and debugging instincts at the right layer.

Do not dump senior-level pressure framing into beginner lessons just to sound advanced.

## Phase-Level Blueprint

### Phase 0: Machine Foundation (`s00`)

Required elements:

- mission and mental model
- visual diagrams
- plain-language analogies
- runnable demonstrations

### Phase 1: Language Foundation (`s01` through `s04`)

Required elements:

- mission
- mental model
- literal or near-literal walkthroughs
- clean runnable code

Avoid:

- premature scale pressure
- advanced security catalogs
- abstract design jargon before concrete examples

### Phase 2: Engineering Core (`s05` through `s08`)

Add:

- trade-off explanations
- failure cases
- safer defaults
- tests and verification surfaces
- performance and maintainability reasoning
- production notes with real-world consequences

### Phase 3: Systems Engineering (`s09` through `s10`)

Add:

- architecture trade-offs
- production notes
- security implications
- deployment patterns
- operational reasoning

### Phase 4: Flagship Project (`s11`)

The flagship applies all prior concepts together:

- integrated project proof
- production deployment thinking
- operational pressure
- security and tenant boundaries
- reliability under failure

## Canonical Lesson Contract

Default lesson shape:

```text
lesson-name/
├── README.md
├── main.go
├── main_test.go
└── _starter/
    └── main.go
```

### Required README Sections

Each lesson README must include these sections in this order:

1. `## Mission`
2. `## Prerequisites`
3. `## Mental Model`
4. `## Visual Model`
5. `## Machine View`
6. `## Run Instructions`
7. `## Code Walkthrough`
8. `## Try It`
9. `## In Production`
10. `## Thinking Questions`
11. `## Next Step`

For exercises:

- replace `## Code Walkthrough` with `## Solution Walkthrough`
- add `## Verification Surface`

### Required Source-File Behavior

Source files should stay readable and runnable.

Use inline comments for:

- non-obvious behavior
- mutation or boundary traps
- subtle runtime implications
- cross-references to earlier or later concepts

Do not turn code headers into the main teaching surface. The README is the primary teaching surface.

## Canonical Milestone Contract

Every section needs proof, not just lessons.

A milestone should usually provide:

- clear README instructions
- a runnable completed solution
- a starter surface when the learner is expected to implement
- tests when behavior should be provable

## Cross-Reference Rules

When a lesson uses a concept not yet formally taught:

- **Forward reference**: explain the borrowed concept briefly and point to the later lesson
- **Backward reference**: name the earlier lesson that established the current idea
- **Sibling reference**: point to the neighboring track when the learner is using, not introducing, the concept

## Adding or Revising Lessons

1. Work inside an existing section.
2. Keep the learner-facing section count at exactly 12.
3. Register the lesson in `curriculum.v2.json`.
4. Update the section README.
5. Ensure `README.md`, `main.go`, tests, starter code, and curriculum metadata agree.
6. Run `go run ./scripts/validate_curriculum.go`.

## Bottom Line

The Go Engineer should feel like one coherent engineering learning system.

The 12 sections are the public spine. The README-first teaching contract is the delivery standard. Future expansion should deepen the sections we have, not fragment the learner path.
