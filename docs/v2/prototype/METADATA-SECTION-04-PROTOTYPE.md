# Prototype Metadata Example: Section 04 Functions And Errors

## Purpose

This document defines the canonical metadata example for the first v2 structural prototype.

It answers the metadata half of issue `#88` and proves that the current schema draft can describe a
prototype section as a connected system rather than as isolated records.

## Prototype Decision

The prototype metadata example should cover:

- one section record
- the key prototype content items that define the section's teaching and milestone arc
- one learning-path layer that shows route logic without duplicating content

That is more useful than a single isolated item because it proves:

- content-type relationships
- prerequisite flow
- milestone outputs
- section-exit navigation

## Canonical Example

```json
{
  "section": {
    "id": "s04-prototype",
    "number": "04",
    "slug": "functions-and-errors",
    "title": "Functions and Errors",
    "phase": "foundations",
    "summary": "Teach learners to shape behavior with functions, model failure explicitly, and cross the boundary from language basics into the first milestone artifact.",
    "status": "prototype",
    "prerequisites": ["s01", "s02", "s03"],
    "path_prefix": "04-functions-and-errors",
    "entry_points": ["FEP.1"],
    "outputs": ["FEP.C1", "FEP.P1"]
  },
  "content_items": [
    {
      "id": "FEP.4",
      "section_id": "s04-prototype",
      "slug": "custom-errors-wrapping-inspection",
      "title": "Custom Errors, Wrapping, and Inspection",
      "type": "lesson",
      "subtype": "pattern",
      "level": "core",
      "verification_mode": "run",
      "estimated_time": 35,
      "summary": "Teach learners to create stable error identities, wrap failures with context, and inspect error chains safely.",
      "objectives": [
        "Define sentinel and typed errors for meaningful failure states",
        "Wrap lower-level errors with contextual information",
        "Use errors.Is and errors.As to inspect an error chain safely"
      ],
      "prerequisites": ["FEP.3"],
      "production_relevance": "Go services rely on wrapped and inspectable errors to preserve control-flow clarity while still exposing enough context for logs, retries, and user-safe handling.",
      "path": "04-functions-and-errors/5-custom-errors-wrapping",
      "run_command": "go run ./04-functions-and-errors/5-custom-errors-wrapping",
      "test_command": "",
      "starter_path": "",
      "next_items": ["FEP.E1"],
      "tags": ["errors", "wrapping", "inspection", "go-idioms"]
    },
    {
      "id": "FEP.E1",
      "section_id": "s04-prototype",
      "slug": "safe-operations-runner",
      "title": "Safe Operations Runner",
      "type": "exercise",
      "subtype": "",
      "level": "core",
      "verification_mode": "mixed",
      "estimated_time": 60,
      "summary": "Guide learners through building a small runner that applies safe operations, wraps context, and reports structured success and failure results across a batch.",
      "objectives": [
        "Implement small operations with explicit value and error contracts",
        "Wrap operation failures with contextual information while preserving inspectable error identity",
        "Process a batch of requests without stopping at the first failure"
      ],
      "prerequisites": ["FEP.1", "FEP.2", "FEP.3", "FEP.4"],
      "production_relevance": "CLI and service code often process small batches, preserve per-item failure context, and continue work while still producing useful summaries.",
      "path": "04-functions-and-errors/6-safe-operations-runner",
      "run_command": "go run ./04-functions-and-errors/6-safe-operations-runner",
      "test_command": "",
      "starter_path": "04-functions-and-errors/6-safe-operations-runner/_starter",
      "next_items": ["FEP.5"],
      "tags": ["exercise", "functions", "errors", "wrapping", "batch-processing"]
    },
    {
      "id": "FEP.5",
      "section_id": "s04-prototype",
      "slug": "defer-panic-failure-boundaries",
      "title": "Defer, Panic, and Failure Boundaries",
      "type": "lesson",
      "subtype": "integration",
      "level": "core",
      "verification_mode": "run",
      "estimated_time": 40,
      "summary": "Teach learners to use defer intentionally, keep panic out of ordinary flow, and recover only at clear safety boundaries.",
      "objectives": [
        "Use defer for cleanup or completion tracking",
        "Explain when panic is inappropriate for ordinary failures",
        "Recover only at a controlled boundary"
      ],
      "prerequisites": ["FEP.E1"],
      "production_relevance": "Reliable Go code uses defer for disciplined cleanup and reserves panic recovery for bounded safety boundaries rather than routine control flow.",
      "path": "04-functions-and-errors/7-defer-panic-failure-boundaries",
      "run_command": "go run ./04-functions-and-errors/7-defer-panic-failure-boundaries",
      "test_command": "",
      "starter_path": "",
      "next_items": ["FEP.C1"],
      "tags": ["defer", "panic", "recover", "failure-boundaries"]
    },
    {
      "id": "FEP.C1",
      "section_id": "s04-prototype",
      "slug": "reliable-batch-processor-checkpoint",
      "title": "Reliable Batch Processor Checkpoint",
      "type": "checkpoint",
      "subtype": "",
      "level": "core",
      "verification_mode": "mixed",
      "estimated_time": 75,
      "summary": "Validate that the learner can build a small batch processor with explicit function contracts, wrapped errors, defer-based cleanup, and panic recovery at a clear boundary.",
      "objectives": [
        "Design a small batch processor that separates per-item work from batch coordination",
        "Use defer for cleanup or completion tracking and convert unexpected panic into controlled error output",
        "Produce a readiness artifact with per-item outcomes and a trustworthy final summary"
      ],
      "prerequisites": ["FEP.1", "FEP.2", "FEP.3", "FEP.4", "FEP.5"],
      "production_relevance": "Reliable worker and batch-style code needs bounded failure behavior, explicit cleanup, and useful summaries so one bad item does not destroy the whole run.",
      "path": "04-functions-and-errors/8-reliable-batch-processor-checkpoint",
      "run_command": "go run ./04-functions-and-errors/8-reliable-batch-processor-checkpoint",
      "test_command": "",
      "starter_path": "",
      "next_items": ["FEP.P1"],
      "tags": ["checkpoint", "functions", "errors", "defer", "panic", "recover", "batch-processing"]
    },
    {
      "id": "FEP.P1",
      "section_id": "s04-prototype",
      "slug": "task-journal-cli",
      "title": "Task Journal CLI Linkage",
      "type": "mini_project",
      "subtype": "",
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
  ],
  "learning_paths": [
    {
      "id": "lp-full",
      "title": "Full Path",
      "audience": "complete beginners",
      "entry_assumptions": "Learner is following sections in order.",
      "recommended_start_items": ["FEP.1"],
      "required_items": ["FEP.4", "FEP.E1", "FEP.5", "FEP.C1", "FEP.P1"],
      "optional_items": [],
      "milestones": ["FEP.C1", "FEP.P1"]
    },
    {
      "id": "lp-bridge",
      "title": "Bridge Path",
      "audience": "experienced programmers who still want proof of readiness",
      "entry_assumptions": "Learner may skim repetition but still needs milestone proof.",
      "recommended_start_items": ["FEP.3"],
      "required_items": ["FEP.4", "FEP.5", "FEP.C1", "FEP.P1"],
      "optional_items": ["FEP.E1"],
      "milestones": ["FEP.C1", "FEP.P1"]
    },
    {
      "id": "lp-targeted",
      "title": "Targeted Path",
      "audience": "learners entering for a specific weak area",
      "entry_assumptions": "Learner completes prerequisite recap before entering Section 04 directly.",
      "recommended_start_items": ["FEP.4"],
      "required_items": ["FEP.C1"],
      "optional_items": ["FEP.E1", "FEP.P1"],
      "milestones": ["FEP.C1"]
    }
  ]
}
```

## What This Example Proves

This prototype example proves that the current schema draft is already rich enough to express:

- section-level identity and outputs
- item-level type and verification differences
- prerequisite flow between lesson, exercise, checkpoint, and project surfaces
- learning-path route logic without duplicating content trees
- section-exit navigation from a mini-project to the next section

## Adjustments Discovered During The Prototype

The prototype surfaces two useful clarifications:

1. section-exit milestone items may point `next_items` at the next section id
2. one metadata example should cover related prototype objects rather than one item in isolation

No additional schema fields are required for the first prototype.

## Success Signal For Issue #88

The metadata half of issue `#88` should be considered successful when maintainers can say:

- this example is realistic enough to guide future tooling work
- the schema feels rich enough without becoming bureaucratic
- the prototype section reads like one connected system
