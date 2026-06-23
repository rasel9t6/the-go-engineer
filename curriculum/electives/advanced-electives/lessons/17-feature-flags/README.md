# Feature flags

## Mission

Understand and apply Feature flags in the context of professional Go software engineering.

## Prerequisites

- elective-16

## Mental Model

A feature flag is a conditional branch controlled by an external configuration system. The flag evaluates to true or false based on context (user ID, request attributes, environment, random percentage). The code path for the new feature is wrapped in an `if flag.Enabled(ctx, "feature-name")` check. The flag can be changed at runtime: enabling the feature for 1% of users, disabling it if errors spike, and removing the flag entirely once the feature is stable.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A feature flag system evaluates flag values based on rules: percentage rollout (enabled for X% of users based on hash(user_id) % 100), user targeting (enabled for specific user IDs or email domains), environment (enabled in staging, disabled in production), and gradual rollout (enabled for increasing percentages over time). Flags are stored in a configuration service (etcd, Consul, LaunchDarkly) and cached locally with a TTL. The local cache avoids per-request network calls. Flag removal is a code cleanup step: after the feature is validated at 100% for a stable period, the flag check and old code path are removed.

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

- Not cleaning up flags after rollout — dead flag branches accumulate, making code harder to read and maintain.
- Using flags for configuration — flags are for features, not for configuration values (database DSN, port number).
- Not monitoring flag usage — without tracking which flags are checked and their values, operator confusion about which features are active.
- Using flags that are checked in too many places — a flag should control a single feature in a single layer, not scattered across handlers, services, and frontend.
- Not rolling back flags — if the new feature causes data corruption, disabling the flag does not undo the corruption (compensating actions needed).

## In Production

Feature flags are used in every continuous delivery pipeline. Companies like Netflix, Facebook, and Google use flags for every feature launch. LaunchDarkly is the leading SaaS feature flag platform. Open-source alternatives: Unleash, Flipt. Go services use flags for gradual rollouts, A/B testing, operational kill switches, and environment-specific behavior.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `elective-18`.
