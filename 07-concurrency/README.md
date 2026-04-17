# 07 Concurrency

## Mission

This stage teaches how Go coordinates work over time, across goroutines, and under cancellation
pressure.

By the end of this stage, a learner should be able to:

- launch and coordinate goroutines deliberately
- use channels and context without losing control of program flow
- reason about timers, scheduling, and cancellation
- apply bounded concurrency patterns with clearer safety instincts

## Stage Map

| Track | Surface | Core Job |
| --- | --- | --- |
| `CON.1` | [concurrency foundations](./01-concurrency) | teach goroutines, channels, context, and scheduling |
| `CON.2` | [concurrency patterns](./02-concurrency-patterns) | teach errgroup, pools, pipelines, and bounded concurrent work |

## Suggested Learning Flow

1. Work through [concurrency foundations](./01-concurrency) first.
2. Move into [concurrency patterns](./02-concurrency-patterns) once goroutine and context behavior feel grounded.
3. Use the exercise-backed capstones as the proof surfaces for this stage.

## Next Step

After this stage, move to [08 Quality & Test](../08-quality-test).
