# Language Basics Track

## Mission

This track teaches the first real Go fundamentals after installation and tooling.

By the end of this track, a learner should be able to:

- declare and update variables confidently
- use constants for fixed values
- model enum-like values with named types and `iota`
- combine those pieces into one small, readable program

## Track Map

| ID | Surface | Why It Matters | Requires |
| --- | --- | --- | --- |
| `LB.1` | [variables](./1-variables) | Introduces typed values, zero values, and the three common declaration styles. | entry |
| `LB.2` | [constants](./2-constants) | Separates immutable values from ordinary variables and clarifies grouped declarations. | `LB.1` |
| `LB.3` | [enums with iota](./3-enums) | Shows how Go models enum-like values without a dedicated enum keyword. | `LB.2` |
| `LB.4` | [application logger](./4-application-logger) | Combines the track into one small milestone exercise with a useful output shape. | `LB.1`, `LB.2`, `LB.3` |

## Engineering Depth

Go is a strictly typed language. Unlike dynamic languages, Go makes the shape of your values
explicit and gives every type a predictable zero value.

That means:

- less guessing about uninitialized state
- clearer compile-time feedback
- simpler reasoning about small programs before control flow becomes more complex

## Live Milestone

[`LB.4 application logger`](./4-application-logger) is the current milestone for this track.

It is intentionally small, but it proves something important: a learner can already combine custom
types, constants, formatting, and defensive checks into one coherent program.

## References

1. [Tour of Go: Variables](https://go.dev/tour/basics/8)
2. [Effective Go: Constants](https://go.dev/doc/effective_go#constants)
3. [Go by Example: Constants](https://gobyexample.com/constants)

## Next Step

After `LB.4`, continue to Section 02: Control Flow.
