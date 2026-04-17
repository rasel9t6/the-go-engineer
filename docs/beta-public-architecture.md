# Source Inventory vs Public 11-Stage Architecture

The Go Engineer now has one canonical learner-facing architecture and one matching set of
root-level stage folders.

This document explains how the current root-stage source surfaces relate to older section
inventory paths that still exist as legacy redirects or contributor history.

## The Short Version

If you are a learner:

- start from [README.md](../README.md)
- then follow [docs/stages/README.md](./stages/README.md)

If you are trying to find files:

- use the current root-stage folders
- use [docs/curriculum/README.md](./curriculum/README.md)

The stage model tells learners **where to go next**.
The folder tree now largely matches that learner-facing model.

## The Public 11 Stages

1. `01 Getting Started`
2. `02 Language Basics`
3. `03 Functions & Errors`
4. `04 Types & Design`
5. `05 Packages & IO`
6. `06 Backend & DB`
7. `07 Concurrency`
8. `08 Quality & Test`
9. `09 Architecture`
10. `10 Production`
11. `11 Flagship`

These 11 stage pages are the public routing surfaces for the curriculum.

## How The Current Source Maps

The current source inventory now maps like this:

- `01-getting-started` -> `01 Getting Started`
- `02-language-basics`, `02-language-basics/03-control-flow`, `02-language-basics/04-data-structures` -> `02 Language Basics`
- `03-functions-errors` -> `03 Functions & Errors`
- `04-types-design`, `05-composition`, `06-strings-and-text` -> `04 Types & Design`
- `05-packages-io/01-modules-and-packages`, `05-packages-io/02-io-and-cli` -> `05 Packages & IO`
- `06-backend-db/01-web-and-database` -> `06 Backend & DB`
- `07-concurrency/01-concurrency`, `07-concurrency/02-concurrency-patterns` -> `07 Concurrency`
- `08-quality-test/01-quality-and-performance` -> `08 Quality & Test`
- `09-architecture/01-package-design`, `09-architecture/02-grpc` -> `09 Architecture`
- `10-production/01-structured-logging`, `10-production/02-graceful-shutdown`, `10-production/03-docker-and-deployment` -> `10 Production`
- `11-flagship/01-enterprise-capstone`, `11-flagship/02-code-generation` -> `11 Flagship`

## Why This Document Still Exists

The move to the 11-stage source surfaces is complete, but older folder names still matter in three
places:

1. historical references in older notes and discussions
2. legacy redirect folders left behind for safer navigation
3. contributor context while curriculum depth continues to grow inside the 11 public stages

## What To Trust For What

### If you are a learner

Use:

- [README.md](../README.md)
- [LEARNING-PATH.md](../LEARNING-PATH.md)
- [docs/stages/README.md](./stages/README.md)

### If you want the physical source layout

Use:

- [docs/curriculum/README.md](./curriculum/README.md)
- the root-stage folders

### If you are a maintainer or contributor

Use:

- the stage pages for public routing truth
- the root-stage folder tree for implementation truth
- `ARCHITECTURE.md` and `CURRICULUM-BLUEPRINT.md` for the longer-range design

## What This Does Not Mean

This transition model does **not** mean:

- learners should navigate by old stage names
- the repo has competing public architectures
- every old planning surface is still public truth
- the older folder names are still the public architecture

## Bottom Line

The 11-stage architecture is the public curriculum, and the root-stage folders are now the
canonical learner-facing source surfaces inside the repo.
