# Prototype Mini-Project: FEP.P1 Task Journal CLI Linkage

## Purpose

This document defines the canonical mini-project for the first v2 structural prototype.

It is the project half of issue `#88` and serves as the reference shape for:

- the first section-exit milestone
- mini-project scope and verification
- project folder layout inside a section
- the bridge from foundations content into later CLI and IO work

## Prototype Decision

The canonical mini-project for the first prototype slice is:

- `FEP.P1 Task Journal CLI Linkage`

Project shape:

- `type`: `mini_project`
- `level`: `core`
- `verification_mode`: `mixed`

## Why `FEP.P1` Is The Right Mini-Project

`FEP.P1` is the right first mini-project because it proves a real curriculum jump without pulling in
infrastructure the learner has not earned yet.

It proves:

- the learner can build a runnable artifact rather than only a lesson-sized example
- the learner can organize a small project with `cmd/` and `internal/`
- the learner can apply explicit error flow to a simple domain workflow
- the section can exit into a believable milestone instead of stopping at a checkpoint

It also creates a clean bridge into later CLI, filesystem, and encoding sections.

## Why It Is Called "CLI Linkage"

This project should prepare the learner for later real CLI work without pretending Section 04
already taught file IO, argument parsing, or persistence.

The "linkage" idea means:

- use a real `cmd/task-journal` entry point
- model command-like actions such as add, complete, and list
- keep the session input simple and foundations-safe
- leave true CLI parsing, storage, and richer IO concerns for later sections

That makes the project forward-looking without becoming dishonest about prerequisites.

## Project Mission

Build a small task journal application that executes a scripted session of journal actions and
prints a trustworthy end-of-run summary.

The learner should demonstrate that they can:

- model a simple task domain
- separate entrypoint, domain logic, and command-style orchestration
- validate and classify expected failures through inspectable errors
- keep the program runnable and understandable as one coherent tool

## Placement In The Curriculum

This mini-project belongs after:

- `FEP.C1 Reliable Batch Processor Checkpoint`

Its job is to answer:

- "Can the learner turn the section's foundations into a small but believable tool?"

And it should point naturally toward:

- Section 09 CLI and IO work

## Canonical Metadata

The mini-project should use this metadata model in the prototype:

```json
{
  "id": "FEP.P1",
  "section_id": "s04-prototype",
  "slug": "task-journal-cli",
  "title": "Task Journal CLI Linkage",
  "type": "mini_project",
  "level": "core",
  "verification_mode": "mixed",
  "estimated_time": 150,
  "summary": "Build a small task journal binary with a real entrypoint, structured domain errors, command-style actions, and a final session summary that prepares learners for later CLI and IO sections.",
  "objectives": [
    "Build a runnable milestone artifact with a real cmd and internal project layout",
    "Apply explicit function and error design to a simple domain workflow",
    "Prepare a foundations project that can later evolve into richer CLI and persistence work"
  ],
  "prerequisites": ["FEP.1", "FEP.2", "FEP.3", "FEP.4", "FEP.5", "FEP.C1"],
  "production_relevance": "Small internal tools often start as bounded command-style programs with clear entrypoints, useful domain errors, and simple package boundaries before they grow into richer CLI workflows.",
  "path": "04-functions-and-errors/9-task-journal-cli",
  "run_command": "go run ./04-functions-and-errors/9-task-journal-cli/cmd/task-journal",
  "test_command": "",
  "starter_path": "",
  "next_items": ["s09"],
  "tags": ["mini-project", "cli", "functions", "errors", "project-layout"]
}
```

## Canonical Project Scope

The learner should build a task journal tool that supports a small scripted session with actions
such as:

- add a task
- complete a task
- list current tasks
- print a final summary

The project should require:

1. a real `main` entrypoint under `cmd/task-journal`
2. a small internal package that owns the journal domain
3. command-style functions that apply journal actions and return explicit errors
4. structured handling for expected failures such as unknown task ids or duplicate task titles
5. a final session summary that distinguishes successful and failed actions

The project should not require:

- persistent storage
- external packages
- advanced flag parsing
- network calls
- multi-module setup

## Expected Folder Shape

The expected prototype folder shape is:

```text
04-functions-and-errors/
  9-task-journal-cli/
    README.md
    cmd/
      task-journal/
        main.go
    internal/
      journal/
        entry.go
        journal.go
        errors.go
      commands/
        runner.go
    testdata/              # optional
```

This project should stay inside the section directory.
It should not move into a top-level `projects/` tree.

## `_starter/` Decision

This mini-project should not include:

- `_starter/`

That is the right choice because the checkpoint already acts as the readiness gate.
By the time the learner reaches the mini-project, they should shape a small artifact with more
independence than a guided exercise allows.

The README should be supportive, but the project itself should not arrive half-built.

## Verification And Success Criteria

This mini-project should use:

- `verification_mode: mixed`

The project should be verified by running the binary and comparing the result against explicit
project criteria in the README.

Minimum success criteria:

- the binary runs from the declared `cmd/task-journal` entrypoint
- valid actions produce the expected journal state
- invalid actions return stable, inspectable errors rather than fragile string-matching behavior
- the run continues through ordinary command failures and still produces a final summary
- the package split is understandable and not overbuilt
- the README explains what later sections would add, such as real argument parsing or storage

## README Expectations

This mini-project should have a learner-facing README with:

1. project mission
2. required features
3. suggested package layout
4. run instructions
5. success criteria
6. extension ideas for later sections

The README should frame the project as a milestone, not as a final polished application.

## Why This Is A Mini-Project Instead Of A Checkpoint

`FEP.P1` differs from the checkpoint because it:

- asks the learner to produce a named runnable artifact
- introduces a real project layout with `cmd/` and `internal/`
- emphasizes deliverable quality and extension potential
- shifts from "prove readiness" to "build the first milestone tool"

Its job is to feel like the learner has built something that can grow.

## Why This Is Not Yet A Phase Capstone

`FEP.P1` still stops well short of capstone scope because it:

- stays in one small domain
- avoids persistence and deployment concerns
- uses a compact package set
- focuses on foundations-plus-Section-04 skills only

That keeps later projects meaningful.

## Schema And Folder-Structure Adjustments Discovered

The prototype surfaces two useful clarifications:

1. planned local path numbering should follow the canonical section order, not the numbering of the
   v1 source material that inspired the slice
2. a section-exit mini-project may use `next_items` to point to the next section id when the next
   meaningful step is a new section rather than one more local item

No new project-specific schema fields are required for the prototype.
The current `type`, `verification_mode`, `path`, `starter_path`, and `next_items` fields are
enough.

## Success Signal For Issue #88

The mini-project half of issue `#88` should be considered successful when maintainers can say:

- this feels larger than a checkpoint but smaller than a capstone
- the project layout is believable and restrained
- the project is honest about current prerequisites
- the project clearly bridges to later CLI and IO work

