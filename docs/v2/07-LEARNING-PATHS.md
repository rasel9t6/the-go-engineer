# V2 Learning Paths

## Purpose

v2 needs explicit learning paths so learners can enter the curriculum honestly without the repo
pretending that every person should take the exact same route.

This document defines how v2 should support different learner backgrounds while still preserving one
canonical curriculum system.

## Core Rule

v2 should have one curriculum with multiple routes through it, not multiple disconnected curricula.

Learning paths should:

- reuse the same canonical sections and content items
- adjust pacing and entry points
- preserve validation at critical transitions
- stay understandable to learners and maintainers

They should not become three separate content trees.

## Why Learning Paths Matter

Without explicit path rules:

- beginners do not know where to start
- experienced programmers skip too much and later get lost
- working Go developers jump straight to advanced material without checking prerequisites
- maintainers write generic guidance that helps no one particularly well

v2 should make the route logic visible instead of burying it in long prose or chat decisions.

## Canonical Path Set

The first v2 draft should support three canonical paths.

### 1. Full Path

Audience:

- Beginner Builder
- any learner who wants the complete guided experience

Route rule:

- follow sections in order
- complete all core lessons, drills, exercises, checkpoints, and required mini-projects

Use this path when:

- the learner is new to programming
- the learner is new to Go and wants maximum scaffolding
- confidence and repetition are more important than speed

### 2. Bridge Path

Audience:

- Transfer Learner

Route rule:

- keep the section order
- allow selective skimming of some foundation lessons
- require the canonical validation items that prove Go-specific understanding

Use this path when:

- the learner already programs in another language
- the learner needs Go mental model correction more than generic programming instruction

### 3. Targeted Path

Audience:

- Working Go Improver

Route rule:

- allow section-based entry
- require explicit prerequisite review and a readiness check
- focus on section outcomes, checkpoints, and production-relevant items

Use this path when:

- the learner already writes some Go
- the learner is trying to improve a targeted area such as concurrency, testing, or architecture

## Path Design Principles

All v2 learning paths should follow these rules.

### One Canonical Object Graph

Paths reference the same canonical sections and content items.
They do not create duplicate lesson copies just because two learners move through them differently.

### Validation Cannot Be Optional Everywhere

Fast paths may skip repetition, but they may not skip proof.

At minimum, every path should preserve:

- section checkpoints
- path-critical exercises
- mini-projects or other milestone artifacts where they carry real curriculum value

### Entry Must Be Honest

If a path assumes prior skill, it should say so directly.
Do not market a fast path as "easy" when it actually depends on unstated knowledge.

### Skipping Should Be Targeted

The Bridge Path and Targeted Path should skip repetition, not fundamentals that the later system
still depends on.

### Paths Must Stay Maintainable

If a path rule requires maintainers to hand-curate dozens of exceptions per section, the design is
too brittle.

## Path Anatomy

Every canonical learning path should define:

- audience
- entry assumptions
- recommended start point
- required content categories
- skippable content categories
- required milestones
- exit signal

## Route Logic By Path

### Full Path Route Logic

Default section flow:

- complete every section in canonical order

Default content expectations:

- all core lessons
- drills when present
- all guided exercises
- all checkpoints
- required mini-projects

Skipping policy:

- do not skip by default

Exit signal:

- learner can move through the entire curriculum with maximum support and no hidden gaps

### Bridge Path Route Logic

Default section flow:

- sections still stay in order
- selected beginner-oriented lessons may be skimmed rather than deeply studied

Default content expectations:

- all section checkpoints
- all path-critical exercises
- milestone mini-projects and any explicitly path-critical projects
- any lesson marked as Go-idiom-critical

Skippable candidates:

- repetitive setup lessons once the learner proves the toolchain basics
- some drills in early sections when the learner already demonstrates command of the concept
- selected beginner examples that do not introduce uniquely Go-specific behavior

Exit signal:

- learner moves quickly without carrying incorrect assumptions about Go's error handling,
  composition, tooling, packages, or concurrency model

### Targeted Path Route Logic

Default section flow:

- enter at a declared section entry point
- review prerequisites before starting
- complete the section's required validation and milestone items

Default content expectations:

- short prerequisite recap or dependency statement
- targeted lessons inside the chosen section
- section checkpoint
- relevant mini-project or integration artifact when it carries the section's teaching value

Skippable candidates:

- earlier sections already mastered by the learner
- repetition-heavy drills outside the chosen focus area

Exit signal:

- learner can use the section's techniques in production-minded work without relying on cargo-cult
  copying

## Required Validation Model

Learning paths should not only define what can be skipped.
They must also define what cannot be skipped.

### Validation Floors

The first v2 draft should use these floors:

- Full Path: complete every checkpoint and required mini-project
- Bridge Path: complete every checkpoint plus path-critical exercises and milestone mini-projects
- Targeted Path: complete the chosen section checkpoint and any required integration artifact

### Path-Critical Items

Some items should be marked as critical for specific paths.

Examples:

- a Go-idiom checkpoint may be critical for the Bridge Path
- a concurrency checkpoint may be critical for the Targeted Path entering Section 12
- a mini-project that proves package boundaries may be critical before advanced architecture work

This gives v2 a more honest fast-track story than "just skip around."

## Path Rules At Section Boundaries

Every section should eventually publish path-aware entry guidance.

At minimum, each section should declare:

- who should start here
- what earlier knowledge is assumed
- what the Full Path must do
- what the Bridge Path may skim
- what the Targeted Path must validate before or during entry
- what output proves readiness to leave the section

## Relationship To Content Types

Learning paths operate over the canonical content type system.

Typical path behavior:

- `lesson`: usually required for the Full Path, selectively skimmable for faster paths
- `drill`: mostly required for Full Path, optional or selective for faster paths
- `exercise`: often required because it proves synthesis
- `checkpoint`: usually mandatory for all paths
- `mini_project`: mandatory when it represents a section or phase milestone
- `capstone`: mandatory when the learner is completing the relevant phase or program span
- `reference`: optional support content unless needed for setup or troubleshooting

## Relationship To The Curriculum Schema

The v2 schema should support learning paths as route definitions, not duplicated content lists.

At minimum, path metadata should be able to express:

- audience
- entry assumptions
- recommended start items
- required items
- optional items
- milestones

Future path metadata may also need:

- path-critical tags
- section entry notes
- validation floor rules

## Docs And Navigation Implications

The public docs system should eventually expose learning paths in a lightweight way.

Recommended navigation surfaces:

- root README: choose your path
- section README: path-aware entry guidance
- dedicated learning path doc: explains the three canonical routes
- checkpoint and project docs: state whether the item is required across all paths or only some

The goal is clarity, not a giant route matrix on every page.

## Example Path Contract

```json
{
  "id": "bridge",
  "title": "Bridge Path",
  "audience": "experienced-programmer-new-to-go",
  "entry_assumptions": [
    "Learner has prior programming experience",
    "Learner can read code and use a command line"
  ],
  "recommended_start_items": ["GS.1", "GS.2", "FE.1"],
  "required_items": ["S01-CHECKPOINT", "S04-MINI-PROJECT", "S08-CHECKPOINT"],
  "optional_items": ["GS.3-DRILL", "CF.1-DRILL"],
  "milestones": ["phase-1-complete", "phase-2-complete"]
}
```

## First Prototype Guidance

The prototype does not need three fully authored learning path documents.

It does need enough structure to prove:

- one section can state path-aware entry rules
- one lesson can identify where it fits in a path
- one exercise or checkpoint can act as a path validation gate
- one mini-project can be marked as a milestone

## Deferred Questions Beyond The First Prototype

These points can stay deferred beyond the first prototype:

- exact thresholds for how many items should be marked as path-critical in a large section
- whether learning paths should live in the same metadata file as sections and content items after
  the prototype stage

## Working Recommendation

For the first v2 implementation:

- keep exactly three canonical paths
- keep the section order stable
- let faster learners skip repetition, not proof
- require milestone mini-projects for the Bridge Path instead of every mini-project
- require Targeted Path entry to include a short prerequisite recap at section entry, not
  necessarily a separate artifact
- make checkpoints the primary path control surface
- keep path rules visible in docs and schema, but do not overbuild a personalized routing engine
