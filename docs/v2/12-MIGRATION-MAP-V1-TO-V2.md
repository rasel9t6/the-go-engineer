# V1 To V2 Migration Map

Status: Draft decision register.

## Migration Principle

v2 should migrate the current repo with intention, not with broad destructive rewrites.

Every major asset in v1 needs one of these decisions:

- keep
- revise
- split
- merge
- remove
- add

## Section Migration Strategy

The current recommendation is to preserve the broad section progression and migrate by improving the
instructional system around it.

| V1 Asset | V2 Decision | Direction |
| -------- | ----------- | --------- |
| Sections 01-04 | Revise | keep sequence, strengthen drills, checkpoints, and first mini-project bridge |
| Sections 05-08 | Revise | keep sequence, standardize lesson spec and design-focused exercises |
| Section 09 | Revise | keep section, formalize filesystem, encoding, and CLI exercise ladder |
| Section 10 | Revise | keep section for first migration wave, tighten boundary between web, database, and service lessons |
| Sections 11-12 | Revise | keep sequence, standardize concurrency checkpoints and production pattern framing |
| Sections 13-14 | Revise | keep sequence, improve architecture/testing/project linkage |
| Section 15 | Revise | preserve code generation core, likely expand toward professional automation and delivery |

## Documentation Migration Strategy

| V1 Doc | V2 Decision | Direction |
| ------ | ----------- | --------- |
| `README.md` | Revise | keep as public landing page, shorten overlap, link to v2 canonical planning and learner docs |
| `LEARNING-PATH.md` | Split | keep current learner guide for v1, move canonical v2 path design into `docs/v2/07-LEARNING-PATHS.md` |
| `docs/curriculum/README.md` | Revise | keep as v1 curriculum reference for now, move v2 planning map into `docs/v2/08-CURRICULUM-MAP.md` |
| `ROADMAP.md` | Revise | keep as public roadmap, reduce deep design detail and link to v2 roadmap docs |
| `CONTRIBUTING.md` | Revise | keep contributor workflow at root, point to v2-specific content standards when they exist |
| `curriculum.json` | Revise later | keep as current source of truth until the v2 schema decision is approved |

## Exercise Migration Strategy

| V1 Exercise State | V2 Decision | Direction |
| ----------------- | ----------- | --------- |
| strong exercise with clear scope | Keep and normalize | add metadata, README, and verification where needed |
| strong idea but weak scaffolding | Revise | add `_starter/`, tests, and clearer requirements |
| oversized "exercise" acting like a project | Split | turn into checkpoint plus mini-project or capstone |
| duplicated or low-signal exercise | Remove or merge | reduce noise and keep the best representative form |

## Folder Structure Strategy

The current recommendation is evolutionary, not revolutionary:

- keep major section directories during the first migration waves
- add standard internal structure inside sections before renaming whole trees
- only introduce major folder changes after the prototype proves the model

This reduces churn for:

- current learners
- existing links
- curriculum metadata
- contributor habits

## Release Line Strategy

The migration plan assumes:

- v1 remains available on `release/v1`
- v2 grows in public on `main`
- alpha releases come from `main`
- `release/v2` is created only when the curriculum reaches feature freeze

## Wave-Based Content Migration

The recommended order is:

1. Define the v2 system
2. Prototype the system
3. Build templates and validator support
4. Migrate foundation sections
5. Migrate design and IO sections
6. Migrate systems, quality, and architecture sections
7. Publish migration guide and stabilize for v2 release

## Content Decision Rules

When deciding what to do with a v1 lesson, use these rules:

- keep when the concept, example, and sequencing are already strong
- revise when the concept is good but the teaching structure is inconsistent
- split when a lesson teaches too many ideas at once
- merge when two lessons are too small to justify separate navigation weight
- remove when the content is redundant or below the v2 quality bar
- add when the learning path has a real gap, not just a nice-to-have topic

## Prototype Requirement

No section should be migrated in-place until these are approved:

- one v2 lesson
- one v2 exercise
- one v2 mini-project
- one v2 curriculum metadata example
- one v2 docs navigation example

That is the minimum proof that the migration system works.
