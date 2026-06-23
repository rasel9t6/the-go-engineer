# Rate limiting

## Mission

Understand and apply Rate limiting in the context of professional Go software engineering.

## Prerequisites

- core-10-17

## Mental Model

Rate limiting is like a entrance turnstile at a stadium. It allows a certain number of people per minute. A token bucket gives each client a bucket of tokens that refill at a fixed rate. Each request consumes one token. When the bucket is empty, requests are denied.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

The token bucket algorithm maintains a bucket with N tokens. Every request removes one token. Tokens are added at a fixed rate (r tokens per second) up to the bucket capacity. This allows bursts of up to N requests while capping the sustained rate. Redis implements this with atomic INCR + EXPIRE.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/18-rate-limiting
go test ./curriculum/modules/10-auth-security/lessons/18-rate-limiting
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Implementing per-endpoint rate limits without global limits — attackers can rotate between endpoints to bypass.
- Using IP-based rate limiting behind a reverse proxy — all requests appear to come from the proxy's IP.
- Rate limiting without a burst allowance — legitimate clients with natural traffic spikes get throttled.
- Applying rate limits after expensive operations — the damage is already done.

## In Production

Rate limiting is universal in production APIs. GitHub (5,000 requests/hour), Twitter, Stripe, and every cloud provider implement rate limiting. Go's x/time/rate package provides a production-ready token bucket implementation.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-19`.
