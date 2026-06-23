# Blue/green deployments

## Mission

Understand and apply Blue/green deployments in the context of professional Go software engineering.

## Prerequisites

- elective-17

## Mental Model

Two identical environments (blue and green) run the application. At any time, one environment serves production traffic (active) and the other is idle or running the previous version (standby). To deploy: deploy the new version to the standby environment, run smoke tests, switch the load balancer to the standby environment (making it active), and keep the old environment as a rollback target.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Blue/green requires: (1) infrastructure that supports two identical environments (separate compute, database, cache), (2) a load balancer or proxy that can switch traffic instantly (HAProxy, Nginx, AWS ALB, Kubernetes Service), (3) database migrations that are backward-compatible (old version can read/write the schema that the new version creates), and (4) a shared session store so users do not lose sessions on switch.

## Run Instructions

```bash
Read the lesson and complete the practice task.
No automated test is required for this lesson.
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Not making database migrations backward-compatible — the old environment cannot read the new schema after migration, making rollback impossible without a reverse migration.
- Skipping smoke tests on the standby environment — the new version is deployed but never tested; the switch puts a broken version into production.
- Not pre-warming caches — the standby environment starts with cold caches; the first requests after switch are slow (cache stampede).
- Using DNS-based traffic switching — DNS changes take minutes to propagate; use a load balancer for instant switching.
- Keeping the old environment for too long — the old environment costs 2x infrastructure; decommission after the new version is verified stable (typically 1-7 days).

## In Production

Blue/green deployments are used by organizations requiring instant rollback and zero-downtime deployments: financial services, healthcare, e-commerce. Cloud providers support blue/green natively: AWS CodeDeploy, Azure Deployment Slots, Google Cloud App Engine. Kubernetes can implement blue/green with two Deployments and label-based Service switching.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-19`.
