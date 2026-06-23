# ABAC and policy checks

## Mission

Understand and apply ABAC and policy checks in the context of professional Go software engineering.

## Prerequisites

- core-10-11

## Mental Model

ABAC is attribute-based authorization: access decisions are made by evaluating policies against user, resource, and environment attributes. Unlike RBAC (which role does the user have?), ABAC asks: does the full context permit this action?

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

ABAC policies are sets of rules that evaluate boolean expressions over attributes. A typical policy says: allow if user.department == resource.department AND resource.classification <= user.clearance_level. The policy engine evaluates all matching rules and applies the first applicable effect (allow/deny).

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/12-abac-and-policy-checks
go test ./curriculum/modules/10-auth-security/lessons/12-abac-and-policy-checks
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Implementing ABAC without a policy engine, leading to convoluted if-else chains in handlers.
- Mixing attribute definitions across layers — defining resource ownership in the handler instead of the policy layer.
- Writing policies that are too permissive because of missing attribute constraints.
- Failing to audit policy evaluation results — ABAC decisions are hard to debug without logs.

## In Production

ABAC is used in AWS IAM (condition statements), Google Cloud IAM, and Kubernetes (admission webhooks). Go microservices use OPA as a sidecar for fine-grained authorization decisions.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-13`.
