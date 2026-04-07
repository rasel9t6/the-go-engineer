# V2 Lesson Template

## Purpose

This document defines the canonical reusable lesson template for v2.

It turns the approved lesson planning and prototype work into a maintainable authoring surface that
contributors can copy without improvising:

- lesson metadata
- code header structure
- lesson anatomy
- README decision rules
- validator expectations

This template is derived primarily from:

- `docs/v2/04-LESSON-SPEC.md`
- `docs/v2/11-CURRICULUM-SCHEMA.md`
- `docs/v2/prototype/LESSON-FEP-4-CUSTOM-ERRORS-WRAPPING.md`
- `CODE-STANDARDS.md`

## When To Use This Template

Use this template when the content item is:

- `type: lesson`

Allowed lesson subtypes in the current v2 draft:

- `concept`
- `pattern`
- `integration`

Do not use this template for:

- drills
- exercises
- checkpoints
- mini-projects
- reference items

## Canonical Lesson Directory Shape

Default lesson layout:

```text
N-lesson-name/
  main.go
  README.md        # optional
  main_test.go     # optional
  testdata/        # optional
```

Use a larger multi-file layout only when the teaching goal genuinely requires it.
Lesson complexity should not grow just because the file layout allows it.

## Metadata Stub

Every v2 lesson should start with a metadata draft like this:

```json
{
  "id": "SX.Y",
  "section_id": "sNN",
  "slug": "lesson-slug",
  "title": "Lesson Title",
  "type": "lesson",
  "subtype": "concept",
  "level": "core",
  "verification_mode": "run",
  "estimated_time": 30,
  "summary": "One-sentence description of the lesson's teaching goal.",
  "objectives": [
    "Primary teaching objective",
    "Optional supporting objective"
  ],
  "prerequisites": ["SX.Z"],
  "production_relevance": "One concrete sentence about where this lesson matters in real Go work.",
  "path": "NN-section-name/N-lesson-name",
  "run_command": "go run ./NN-section-name/N-lesson-name",
  "test_command": "",
  "starter_path": "",
  "next_items": ["SX.Y+1"],
  "tags": ["topic-a", "topic-b"]
}
```

## Metadata Field Notes

- `type` must always be `lesson`
- `subtype` should be one of `concept`, `pattern`, or `integration`
- `level` should usually be `foundation`, `core`, `stretch`, or `production`
- `verification_mode` should be `run` unless a test-driven lesson shape is clearly better
- `objectives` should stay small and focused
- `next_items` should point to the most meaningful immediate next step, not every possible path

## Lesson Header Template

Every lesson entry file should begin with a header shaped like this:

```go
// ============================================================================
// Section NN: Section Name - Lesson Title
// Level: Foundation | Core | Stretch | Production
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Primary concept or pattern
//   - Important supporting concept
//   - One meaningful tradeoff or habit
//
// ENGINEERING DEPTH:
//   Why this matters in production Go. What failure mode, habit, or workflow
//   does this lesson prepare the learner for?
//
// RUN: go run ./NN-section-name/N-lesson-name
// ============================================================================
```

Keep this header consistent with `CODE-STANDARDS.md`, but prefer v2 field language such as
`Foundation | Core | Stretch | Production` over the older `Beginner | Intermediate | Advanced`
labels when drafting new v2 content.

## Canonical Lesson Skeleton

Use this as the default code-first lesson shape:

```go
package main

import "fmt"

// ============================================================================
// Section NN: Section Name - Lesson Title
// Level: Core
// ============================================================================
//
// WHAT YOU'LL LEARN:
//   - Primary teaching objective
//   - Supporting concept
//
// ENGINEERING DEPTH:
//   One practical explanation of where this lesson matters in real Go work.
//
// RUN: go run ./NN-section-name/N-lesson-name
// ============================================================================

func main() {
	// Framing:
	// Explain what the learner is about to see and why it matters here.

	// Core example:
	// Show one bounded example that expresses the main idea clearly.

	// Explanation:
	// Use comments around the example to explain tradeoffs, gotchas, and why
	// the code is shaped this way.

	// Production relevance:
	// Connect the example to a real workflow, boundary, or failure mode.

	// Exit ramp:
	// Point the learner to the next lesson, drill, exercise, or checkpoint.

	fmt.Println("replace with lesson output")
}
```

## Required Lesson Anatomy Checklist

Every v2 lesson should clearly cover these five layers:

1. Framing
2. Core example
3. Explanation
4. Production relevance
5. Exit ramp

If one of those layers is missing, the lesson is probably not ready.

## Subtype Guidance

### Concept Lesson

Use when:

- introducing a new language or design idea
- teaching one clear mental model

Default shape:

- one runnable example
- tight explanation
- low architecture overhead

### Pattern Lesson

Use when:

- teaching a reusable engineering habit
- comparing good and bad approaches
- showing one bounded layered example

Default shape:

- one strong pattern example
- explicit tradeoffs
- clear production boundary

### Integration Lesson

Use when:

- combining several prior ideas in a controlled example
- preparing the learner for a checkpoint or mini-project

Default shape:

- one synthesis example
- explicit relationship to prior lessons
- clear warning against overexpansion

## README Decision Rules

The default v2 lesson should be:

- code-first

Add `README.md` only when at least one of these is true:

- runtime setup is non-trivial
- the lesson uses multiple files or commands
- a table, diagram, or structured comparison materially helps understanding
- troubleshooting guidance is necessary beyond inline comments

Do not add a README just to satisfy ceremony.

## Test Decision Rules

Add `main_test.go` when:

- the behavior being taught is naturally testable
- the lesson also teaches validation or correctness checks
- a test clarifies the lesson better than more prose would

Do not add tests to every lesson by default.

## Scope Control Rules

Split the lesson when:

- more than one major new concept is being introduced
- the example is large enough that the core idea is hard to see
- the explanation depends on several unrelated digressions

Keep the lesson compact by preferring:

- one strong example
- one primary objective
- one explicit next step

## Copy-Ready Authoring Checklist

Before calling a lesson draft complete, confirm:

- metadata matches the actual lesson
- the lesson has one clear primary objective
- the run or test command is real
- the production relevance is explicit
- the next step is named clearly in `next_items`
- README support is intentional, not automatic

## Validator Notes

The validator should eventually enforce these lesson-specific checks:

- `type` is `lesson`
- `subtype` uses an allowed lesson subtype
- `path` exists
- either `run_command` or `test_command` points to a real target
- `next_items` resolves
- declared README support is reflected by real files once the schema carries that signal

The validator should not try to judge:

- whether the example is pedagogically strong
- whether the explanation is vivid enough
- whether the lesson should have been split earlier

Those remain reviewer judgments.

## Relationship To Existing Public Standards

This template does not replace the public standards yet.

For now:

- `CODE-STANDARDS.md` remains the public baseline in `main`
- this template is the v2 planning/prototype-derived maintainer template on `planning/v2`

When learner-facing v2 implementation begins, the repo can decide how to merge the two cleanly.

## Success Signal

This template is working when a maintainer can copy it and produce a new v2 lesson without
guessing:

- the metadata shape
- the file header
- the lesson anatomy
- whether a README is needed
- what validator expectations the lesson should satisfy

