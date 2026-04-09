# DS.6 Contact Manager

## Mission

Build a small in-memory contact manager that turns the Section 03 data-structure lessons into one
coherent runnable exercise.

This is the Section 03 milestone.
It is where slices, maps, pointers, and simple struct modeling come together.

## Prerequisites

Complete these first:

- `DS.1` arrays
- `DS.2` slices
- `DS.3` maps
- `DS.4` pointers
- `DS.5` advanced slicing

## What You Will Build

Implement a small contact manager that:

1. defines a `Contact` struct
2. stores contacts in a slice
3. uses a map for efficient name lookup
4. updates contacts through pointers or other mutation-aware logic
5. prints a small demonstration flow in `main()`

## Files

- [main.go](./main.go): complete solution with teaching comments
- [_starter/main.go](./_starter/main.go): starter file with TODOs

## Run Instructions

Run the completed solution:

```bash
go run ./03-data-structures/6-contact-manager
```

Run the starter:

```bash
go run ./03-data-structures/6-contact-manager/_starter
```

## Success Criteria

Your finished solution should:

- define a clear contact data model
- add and retrieve contacts safely
- show at least one update that persists correctly
- keep the flow simple enough that the data-structure choices are visible

## Common Failure Modes

- appending to a slice without returning or reusing the updated slice
- copying data when the operation really needs mutation
- using pointers everywhere even when plain values would stay simpler
- hiding the exercise behind too much package structure

## Next Step

After this milestone, continue to Section 04: Functions and Errors.
