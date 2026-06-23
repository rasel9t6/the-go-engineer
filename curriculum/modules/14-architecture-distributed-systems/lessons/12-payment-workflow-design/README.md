# Payment workflow design

## Mission

Understand and apply Payment workflow design in the context of professional Go software engineering.

## Prerequisites

- core-14-11

## Mental Model

A payment workflow is a state machine with retries and compensating actions. Every state transition must be: (1) persisted before execution (write the intent, then act), (2) idempotent (same transition applied twice is harmless), (3) reversible (if capture fails, void the authorization). The state machine lives in the database, not in memory — if the service restarts, it resumes from the last persisted state. Each step forward has a corresponding step backward (void, refund).

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Payment gateways (Stripe, Adyen, Square) provide REST APIs for payment processing. The standard flow: POST /v1/payment_intents (create with amount), POST /v1/payment_intents/{id}/confirm (authorize), POST /v1/payment_intents/{id}/capture (capture). Each call supports an Idempotency-Key header: if the same key is sent within 24 hours, the gateway returns the same response without executing the operation again. In Go, the gateway client (e.g., stripe-go) generates the idempotency key automatically if not provided. The payment state machine is stored in the application database: the PaymentIntent record tracks the current status, gateway reference IDs, and idempotency keys. The workflow is driven by a scheduler or queue: consumer goroutines check for PaymentIntents in 'authorized' status and attempt capture. If the service restarts, it scans for incomplete PaymentIntents and resumes processing.

## Run Instructions

```bash
go run ./curriculum/modules/14-architecture-distributed-systems/lessons/12-payment-workflow-design
go test ./curriculum/modules/14-architecture-distributed-systems/lessons/12-payment-workflow-design
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Processing payments synchronously in the HTTP handler — the request blocks for the duration of the external payment gateway call (500ms-5s), occupying a goroutine and a database connection. If the gateway is slow, all handler goroutines are blocked and the service stops accepting requests. Payments should be processed asynchronously via a queue.
- Not handling partial failures in multi-step payment workflows — a payment flow involves: authorize, capture, settle, notify. If the settlement succeeds but the notification fails, the payment is settled but the user never receives a receipt. Each step must have its own error handling, retry logic, and compensating action.
- Storing raw payment gateway responses without mapping to domain events — storing Stripe's full JSON response as raw JSONB means the application logic depends on Stripe's API shape. If Stripe adds or removes fields, the parsing breaks. Map gateway responses to domain events (PaymentAuthorized, PaymentCaptured, PaymentFailed) at the adapter boundary.

## In Production

Payment workflows are the most critical path in any e-commerce or SaaS system. Stripe processes billions of dollars/day through their API. Every payment integration follows the same pattern: authorize (hold the funds), capture (charge the card), settle (transfer to merchant account). Failure at any step requires compensating actions (authorization expires after 7 days, capture can be voided, settled amounts must be refunded). Go production services use this pattern for: subscription billing, one-time purchases, marketplace payments, and invoicing.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-14-13`.
