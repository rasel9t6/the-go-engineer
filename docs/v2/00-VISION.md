# V2 Vision

## Thesis

v2 should turn The Go Engineer from a strong collection of lessons into a coherent Go engineering
training system.

The repo already proves there is real content value in:

- practical examples
- project-driven learning
- engineering depth
- a full-stack view of Go

What v2 needs is stronger system design around that content.

## Product Definition

The Go Engineer v2 is a repo-first, open-source Go engineering curriculum that teaches learners from
first program to production-minded systems work through:

- layered lessons
- drills and guided exercises
- section checkpoints
- mini-projects
- phase capstones
- explicit learning paths
- validated curriculum metadata

## What v1 Already Does Well

v2 should preserve these strengths from v1:

- strong topic breadth across language, runtime, IO, web, concurrency, quality, and architecture
- practical examples instead of abstract lecture notes
- visible engineering ambition, especially in the later sections
- curriculum metadata and dependency thinking already present in `curriculum.json`
- a contributor culture that values comments, clarity, and hands-on learning

## Problems v2 Must Solve

v2 exists because v1 has grown faster than its governing system.

The main problems are:

- lesson structure is not yet enforced as a curriculum contract
- exercises exist, but they do not yet form a complete exercise bank system
- docs are useful, but the root README, learning path, roadmap, and curriculum docs overlap
- learner progression is present, but not yet expressed through standard checkpoints and projects
- contributor guidance exists, but the repo still allows drift between content, docs, and metadata

## Target Learners

v2 is designed for three learner groups:

### 1. Complete Beginners

Need slow ramp-up, strong explanation, and confidence-building exercises.

### 2. Developers New To Go

Need Go-specific mental models, idioms, and practice without repeating generic programming theory.

### 3. Working Go Developers

Need targeted depth in concurrency, testing, performance, architecture, and production practice.

## Strategic Position

The honest position for v2 is:

- not an LMS
- not a video course replacement
- not a shallow interview-prep repo
- not a random collection of snippets

It should be a durable, opinionated engineering curriculum that happens to live in Git.

## Non-Goals

v2 should explicitly avoid:

- rewriting every section before shipping any improvement
- building a full learning platform before the curriculum model is proven
- changing names and folder structure just to look "new"
- turning every lesson into a giant project
- adding advanced topics that weaken the beginner-to-engineer arc

## Success Criteria

v2 is successful if it delivers the following:

1. A new contributor can understand the curriculum system from the docs alone.
2. A learner can see where lessons, exercises, checkpoints, and projects fit together.
3. Every section has a predictable instructional shape.
4. The curriculum validator can verify more than simple path existence.
5. Stable v1 users can continue learning while v2 grows in public.
6. The first v2 alpha release already feels more coherent than late-stage v1.

## Release Philosophy

v2 should be shipped in public, section by section.

That means:

- develop on `main`
- support stable users on `release/v1`
- publish `v2.0.0-alpha.N` from meaningful checkpoints on `main`
- create `release/v2` only when the curriculum enters feature freeze

This keeps the migration honest and visible.

## Decision Statement

We are not choosing between "keep v1 forever" and "rewrite everything at once."
We are choosing a third path:

build v2 deliberately, validate the model early, and migrate the curriculum in controlled waves.
