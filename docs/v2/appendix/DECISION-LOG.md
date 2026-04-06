# V2 Decision Log

Use this file to record planning decisions that affect v2 structure, sequencing, migration, or
release strategy.

## How To Use This Log

Each decision should capture:

- date
- decision summary
- why the decision was made
- alternatives considered
- follow-up actions

## Open

### 2026-04-06: Should v2 keep the exact v1 section numbering and directory layout?

- Status: Open
- Why it matters: changing numbering early increases churn for learners, links, and migration work
- Alternatives:
  - keep the current section numbering for the first migration waves
  - redesign numbering and directory layout before any migration
- Follow-up: decide after the prototype and folder-structure draft exist

### 2026-04-06: Should `curriculum.json` be extended or replaced for v2?

- Status: Open
- Why it matters: exercise metadata, project metadata, and learning-path metadata likely exceed the
  current schema
- Alternatives:
  - extend `curriculum.json`
  - add a v2 schema next to the current file and migrate later
- Follow-up: decide after the curriculum schema draft and prototype metadata example are reviewed

### 2026-04-06: How much assessment structure should exist before any platform work is considered?

- Status: Open
- Why it matters: v2 should stay repo-first and avoid platform sprawl
- Alternatives:
  - keep assessment entirely repo-native
  - add richer progress and assessment layers later after the curriculum model is stable
- Follow-up: revisit after the first alpha wave
