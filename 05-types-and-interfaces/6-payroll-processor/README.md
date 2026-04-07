# TI.6 Payroll Processor Project

## Mission

Build a small payroll system that proves multiple concrete employee types can share one behavior
contract through interfaces instead of inheritance.

This exercise is the Section 05 milestone.
It is where structs, methods, interfaces, and a small generic helper come together in one runnable
artifact with tests.

## Prerequisites

Complete these first:

- `TI.1` structs
- `TI.2` methods
- `TI.3` interfaces
- `TI.4` Stringer
- `TI.5` generics

## What You Will Build

Implement a small payroll processor that:

1. defines a `Payable` interface with behavior contracts
2. models multiple concrete employee types with structs
3. attaches pay-calculation behavior through methods
4. processes a mixed `[]Payable` collection polymorphically
5. includes one small generic helper for reusable summary/report output
6. verifies core pay calculations with tests

## Files

- [main.go](./main.go): complete solution with teaching comments
- [payroll_test.go](./payroll_test.go): tests for the pay-calculation contract
- [_starter/main.go](./_starter/main.go): starter file with TODOs and requirements

## Run Instructions

Run the completed solution:

```bash
go run ./05-types-and-interfaces/6-payroll-processor
```

Run the tests:

```bash
go test ./05-types-and-interfaces/6-payroll-processor
```

Run the starter:

```bash
go run ./05-types-and-interfaces/6-payroll-processor/_starter
```

## Success Criteria

Your finished solution should:

- use interfaces to describe the payroll behavior boundary
- allow multiple employee types to satisfy the same contract cleanly
- keep receiver choices intentional and easy to explain
- include one small generic helper without turning the exercise into a generics showcase
- pass the provided tests

## Common Failure Modes

- building separate payroll logic for each employee type instead of using one interface
- using pointer receivers or value receivers without being able to explain why
- treating generics as the main point of the exercise instead of a small reuse helper
- tightly coupling `ProcessPayroll` to one concrete employee struct

## Next Step

After you complete this exercise, move to [TI.7 advanced generics](../7-advanced-generics) for a
stretch lesson, or continue to [Section 06](../../06-composition) if you are ready to move on.
