# RBAC

## Mission

Understand and apply RBAC in the context of professional Go software engineering.

## Prerequisites

- core-10-10

## Mental Model

RBAC is a role-based permission matrix. Users are assigned roles, roles are assigned permissions, and permissions grant access to actions. It is a many-to-many relationship: one user can have multiple roles, one role can grant multiple permissions.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

RBAC has three core entities: users, roles, and permissions. Permissions are typically expressed as action:resource pairs (e.g., post:create, post:delete). A permission check iterates the user's roles and looks up the intersection of role permissions with the required permission.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/11-rbac
go test ./curriculum/modules/10-auth-security/lessons/11-rbac
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Hardcoding role checks in business logic instead of using a policy layer — changes require code deploys.
- Using only roles without resource-level scoping — an admin role should not grant access to all resources by default.
- Storing role assignments in application code rather than a database.
- Implementing RBAC without default-deny — missing role checks implicitly allow access.

## In Production

RBAC is the standard authorization model for enterprise applications. Kubernetes, AWS IAM, and GitHub all use role-based access control. Go services implement RBAC via Casbin or custom middleware.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-12`.
