# Threat modeling

## Mission

Understand and apply Threat modeling in the context of professional Go software engineering.

## Prerequisites

- core-09-28

## Mental Model

Threat modeling is an architectural review: at every trust boundary, ask 'what could go wrong?' Each answer becomes a security requirement.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Threat modeling in Go follows STRIDE/LINDDUN like any language but benefits from Go's explicit error handling. Go's io.Reader/Writer interfaces make data flow tracing straightforward.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/01-threat-modeling
go test ./curriculum/modules/10-auth-security/lessons/01-threat-modeling
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Drawing threat models after an incident instead of during design — post-hoc models miss the assumptions that led to the breach.
- Modeling only external threats and ignoring insider risk — most breaches involve compromised credentials.
- Stopping at diagram creation without tracing trust boundaries, data flows, and attack surfaces through the codebase.

## In Production

Production Go services at Google, Stripe, and Docker use threat modeling as a standard design-phase activity. Security review gates in CI/CD require an up-to-date threat model.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-02`.
