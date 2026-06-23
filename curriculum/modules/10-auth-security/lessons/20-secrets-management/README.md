# Secrets management

## Mission

Understand and apply Secrets management in the context of professional Go software engineering.

## Prerequisites

- core-10-19

## Mental Model

Secrets are the keys to your kingdom — they unlock databases, third-party APIs, and encryption. A secret manager is a secure vault with access controls, audit logs, and automatic rotation. The application asks the vault for the key each time it needs it, not before.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Secret managers store secrets encrypted at rest and in transit. Access is controlled by IAM policies, and every read is audited. Applications authenticate using a short-lived token (e.g., Vault's Kubernetes auth or AWS IAM roles). Secrets are typically fetched once at startup and cached in memory for the application's lifetime.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/20-secrets-management
go test ./curriculum/modules/10-auth-security/lessons/20-secrets-management
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Hardcoding secrets in source code that gets committed to version control.
- Storing secrets in environment variables that are logged or exposed in error pages.
- Using a single secret for multiple purposes (JWT signing, DB encryption, API keys).
- Not rotating secrets regularly — a leaked secret remains valid indefinitely.

## In Production

Every production system uses secrets management. HashiCorp Vault, AWS Secrets Manager, GCP Secret Manager, and Kubernetes Secrets are the standard tools. Go applications integrate with these using official SDKs.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-21`.
