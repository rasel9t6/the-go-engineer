# V2 Curriculum Philosophy

## Core Position

The Go Engineer teaches Go as an engineering practice, not just as a language syntax set.

v2 should preserve the practical style of v1 while making the learning system explicit.

## Teaching Principles

### 1. Learn By Building

Every important concept should appear in executable code, not only in prose.

### 2. Explain The Runtime Model Early

Go becomes easier when learners understand memory, values, references, concurrency, and errors as
behavioral models rather than isolated keywords.

### 3. Production Relevance Matters

Every section should answer:

- Why does this matter in real Go systems?
- What breaks when it is misunderstood?
- What habit should the learner carry into production work?

### 4. Repetition Is A Feature

v2 should intentionally revisit key ideas at higher levels:

- error handling
- interfaces and abstraction
- IO boundaries
- concurrency safety
- testability
- package design

### 5. Synthesis Is Required

Knowing isolated lessons is not enough.
The curriculum must repeatedly force learners to combine concepts through exercises and projects.

## Instructional Stack

Each phase of the curriculum should use the same layered model:

1. Lesson
2. Drill
3. Guided exercise
4. Section checkpoint
5. Mini-project
6. Phase capstone

Not every section needs all six layers in the same size, but every section should contain a visible
practice and synthesis path.

## Lesson Design Rules

Every v2 lesson should:

- teach one primary idea and a small number of supporting ideas
- include a clear "why this matters" framing
- show runnable code
- point back to prerequisites and forward to the next step
- avoid being so small that it becomes trivia
- avoid being so large that it becomes a hidden mini-course

## Exercise Design Rules

Exercises should not only ask learners to copy a pattern.
They should require:

- combination of prior concepts
- decisions under constraints
- debugging or validation thinking
- naming and structure choices
- reflection on tradeoffs

## Project Design Rules

Projects in v2 must do real curriculum work.
Each project should exist because it consolidates a part of the learning arc.

Good reasons for a project:

- it forces package boundaries
- it forces persistence or IO
- it introduces testing surfaces
- it creates a realistic concurrency or HTTP boundary
- it serves as a bridge to a later capstone

Bad reasons for a project:

- it is only flashy
- it duplicates a previous exercise with more lines
- it introduces too many unrelated concepts at once

## Difficulty Model

v2 should use an explicit difficulty ladder:

- Foundation
- Core
- Stretch
- Production

This is not a prestige label. It is a pacing tool.

## Learning Path Philosophy

The curriculum should support different learners without becoming fragmented:

- beginners follow the full path
- experienced programmers use a reduced bridge path
- working Go developers can use section-targeted paths

All paths should still point back to the same canonical content objects.

## Quality Bar

v2 is only coherent if each section meets the same basic standard:

- consistent metadata
- predictable lesson and exercise shape
- clear navigation
- explicit outcomes
- validated runnable paths
- visible production relevance

If a section does not meet that bar, it is not finished, even if the code runs.
