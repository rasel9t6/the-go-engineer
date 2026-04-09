# CF.4 Pricing Calculator

## Mission

Build a small pricing calculator that turns the control-flow basics into one readable,
runnable program.

This is the Section 02 milestone.
It is where loops, `if` branches, `switch`-style thinking, map lookups, and string suffix handling
come together.

## Prerequisites

Complete these first:

- `CF.1` for loop
- `CF.2` if / else
- `CF.3` switch

## What You Will Build

Implement a small order processor that:

1. stores product prices in a map
2. calculates a regular item price safely
3. detects `_SALE` items and applies a discount
4. returns `(price, found)` for each lookup
5. iterates over an order and computes a subtotal

## Files

- [main.go](./main.go): complete solution with teaching comments
- [_starter/main.go](./_starter/main.go): starter file with TODOs
- [calculator_test.go](./calculator_test.go): focused tests for `calculateItemPrice`

## Run Instructions

Run the completed solution:

```bash
go run ./02-control-flow/4-pricing-calculator
```

Run the starter:

```bash
go run ./02-control-flow/4-pricing-calculator/_starter
```

Run the tests:

```bash
go test ./02-control-flow/4-pricing-calculator
```

## Success Criteria

Your finished solution should:

- use a package-level price map
- return `(price, found)` from `calculateItemPrice`
- handle `_SALE` items correctly
- skip unknown items safely
- compute the subtotal from the valid items only

## Common Failure Modes

- assuming a map lookup always succeeds
- forgetting to strip `_SALE` before the second lookup
- mixing printing logic and subtotal logic so heavily that the flow becomes hard to read
- treating missing items as a zero-price success

## Next Step

After this milestone, continue to Section 03: Data Structures.
