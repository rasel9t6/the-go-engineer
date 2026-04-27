# Roadmap

This document tracks the current post-v2.1 state of The Go Engineer.

If this file and [ARCHITECTURE.md](./ARCHITECTURE.md) disagree on structure, `ARCHITECTURE.md` wins.

## Current Status

The v2.1 stable release is shipped.

- the 12-section architecture is live and locked
- root source folders align to the public section map
- `curriculum.v2.json` is the active curriculum registry
- the Go validator is the required curriculum health check
- migration work is complete

The current major stream is **post-v2.1 implementation work on `main`**, with emphasis on Opslane, polish, tests, and repo hygiene.

## Branch Model

- `main`: post-v2.1 implementation line
- `release/v1`: stable v1 maintenance line
- `release/v2`: stable v2.1.x maintenance line

## Stable Snapshot

| Area | Status | Notes |
| --- | --- | --- |
| Public architecture | Complete | 12 sections aligned across root folders and metadata |
| Curriculum metadata | Complete | `curriculum.v2.json` is current and validated |
| Learner-facing section roots | Complete | section entry points exist from `s00` to `s11` |
| Validator | Complete | single Go validator is the required curriculum health check |
| Architecture migration | Complete | no remaining architecture migration blocker |
| Workflow docs | Active polish | keep branch, commit, review, and release docs aligned |
| Opslane flagship | Active implementation | deepen integrated system behavior on `main` |

## Post-Release Quality Focus

Post-release work should concentrate on:

1. learner-path polish and consistency
2. validator strictness and repo hygiene
3. missing test surfaces where behavior should be provable
4. documentation cleanup and release metadata
5. flagship depth and integration confidence

## Section Status

| Section | Status | Current Focus |
| --- | --- | --- |
| s00 How Computers Work | Stable | polish diagrams and machine explanations where useful |
| s01 Getting Started | Stable | keep setup surfaces crisp and trustworthy |
| s02 Language Basics | Stable | preserve zero-magic flow across language, control flow, and data structures |
| s03 Functions & Errors | Stable | tighten proof surfaces and orchestration clarity |
| s04 Types & Design | Stable | keep canonical path coherent while stretch content stays clearly optional |
| s05 Packages, I/O & CLI | Stable | strengthen milestone polish and cross-track navigation |
| s06 Backend, APIs & Databases | Stable | keep HTTP, API, and DB tracks aligned and explicit |
| s07 Concurrency | Stable | preserve clear path from fundamentals to patterns |
| s08 Quality & Testing | Stable | keep testing and profiling proof surfaces honest |
| s09 Architecture & Security | Stable | sharpen trade-off teaching and security progression |
| s10 Production Operations | Stable | keep config, observability, deployment, and code generation coherent |
| s11 Opslane Flagship | Active implementation | deepen the integrated production-shaped capstone |

## Version Plan

| Version | Target | Criteria |
| --- | --- | --- |
| v2.1.0 | released | stable curriculum release is published |
| v2.1.x | maintenance | stable fixes and low-risk corrections on `release/v2` |
| post-v2.1 | current on `main` | flagship implementation and deeper engineering content |

## Success Criteria

After any post-release change, we should still be able to say:

- public docs, metadata, and validator agree
- stable maintenance stays bounded and low-noise
- flagship implementation advances without destabilizing the release line
- engineering gains come from integrated system work, tests, and polish, not architecture churn
