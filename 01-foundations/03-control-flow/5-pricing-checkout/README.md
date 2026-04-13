# CF.5 Pricing Checkout

## Mission

Build one small checkout flow that turns the earlier control-flow lessons into a single readable,
runnable program.

This is the section milestone.

## Why This Milestone Exists

The earlier lessons each taught one control-flow idea in isolation.

This milestone asks the learner to combine them:

- `if` for ordinary decisions
- `for` for repeated work
- `switch` for readable multi-branch rules
- `continue` for skipping invalid items

## Zero-Magic Boundary

This milestone uses a small list of item codes as a preview tool so the learner has something to
process in a loop.

That list is not the topic here.
Formal collection mechanics belong to the next section.

This milestone also avoids helper functions on purpose.
Section 05 is where reusable logic is taught properly.

## What You Will Build

Implement a checkout flow that:

1. loops over a small cart of item codes
2. uses `switch` to assign a base price
3. detects `_SALE` items with an `if` check
4. skips unknown items with `continue`
5. computes a running subtotal

## Files

- [main.go](./main.go): complete runnable solution
- [_starter/main.go](./_starter/main.go): starter file with TODOs

## Run Instructions

```bash
go run ./01-foundations/03-control-flow/5-pricing-checkout
```

```bash
go run ./01-foundations/03-control-flow/5-pricing-checkout/_starter
```

## Solution Walkthrough

### `cart := []string{ ... }`

The cart is a small preview list of item codes.
The important control-flow idea is not the list itself.
The important idea is that the loop now has several inputs to process.

### `for _, item := range cart`

This repeats the same decision process for each cart entry.

### `isSale := strings.HasSuffix(item, "_SALE")`

This creates a yes/no decision the program can branch on.

### `baseCode := strings.TrimSuffix(item, "_SALE")`

If the item is a sale item, the code strips the suffix before pricing it.
That keeps the later `switch` focused on the real product code.

### `switch baseCode { ... }`

This is the pricing rule engine.
Each known code maps to a base price.

### `if price == 0 { ... continue }`

Unknown items are skipped safely.
The loop keeps running, but the subtotal is not polluted by invalid entries.

### `if isSale { price = price * 0.80 }`

This applies the discount only after the base price has been found.

### `subtotal += price`

This is the running-total pattern.
Each valid item contributes to the final answer.

## Try It

1. Add another known item code to the cart.
2. Add one more `_SALE` item and confirm the discount still applies.
3. Put an unknown item in the middle of the cart and watch the loop skip it cleanly.
4. Change the discount from `20%` to `10%`.

## Common Failure Modes

- trying to price the `_SALE` code before stripping the suffix
- forgetting to skip unknown items
- applying the discount before a valid base price exists
- mixing too many future abstractions into a control-flow milestone

## Success Criteria

Your program is successful when it:

- processes the whole cart from top to bottom
- prices known items correctly
- discounts sale items correctly
- skips unknown items safely
- prints a final subtotal

## Next Step

After this milestone, continue to `04-data-structures`.
That is where list and lookup behavior become the real topic instead of a preview tool.
