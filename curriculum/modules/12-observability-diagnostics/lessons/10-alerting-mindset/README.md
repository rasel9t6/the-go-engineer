# Alerting mindset

## Mission

Understand and apply Alerting mindset in the context of professional Go software engineering.

## Prerequisites

- core-12-09

## Mental Model

Alerting is a smoke detector, not a thermostat. A thermostat constantly adjusts (low-level monitoring). A smoke detector only makes noise when there is a fire (user-impacting problem). If the smoke detector goes off every time someone cooks (every deployment causes a brief latency spike), the occupants disable it (alert fatigue). The correct design: the smoke detector only goes off for actual fires — situations that require immediate human action. Everything else is a dashboard or log.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Alertmanager receives alerts from Prometheus via the Alertmanager API. Each alert has: labels (severity, service, team), annotations (summary, description, runbook_url), and a status (firing or resolved). Alertmanager applies grouping: alerts with the same group_by labels are sent as a single notification (e.g., group all alerts for service='api' into one Slack message). Inhibitions suppress low-severity alerts when a high-severity alert is firing (e.g., if the database is down, suppress all application-level alerts that depend on it). Silences mute alerts for a specified duration (used for maintenance). The notification template (default or custom Go template) formats the alert into a message for the target (Slack webhook, PagerDuty API, email). Prometheus alert rules are evaluated as part of PromQL query evaluation — the rule expression is executed against the TSDB at the evaluation interval, and the result (truth value of the expression) determines the alert state.

## Run Instructions

```bash
go run ./curriculum/modules/12-observability-diagnostics/lessons/10-alerting-mindset
go test ./curriculum/modules/12-observability-diagnostics/lessons/10-alerting-mindset
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Alerting on every symptom — creating an alert for every metric that can go wrong (high CPU, high memory, slow query, error count > 0) produces thousands of alerts that operators ignore. An alert should represent a user-impacting condition that requires human action within a specific time window. If the team ignores an alert, it should be deleted or downgraded to a warning.
- Setting static thresholds without historical context — alerting when p99 latency > 500ms may fire constantly for a service that normally runs at 480ms p99. The threshold must be based on the service's actual performance distribution, not an arbitrary number. Use dynamic thresholds (baseline + 3σ) or set the threshold at the SLO boundary plus a safety margin.
- Not defining a runbook before creating the alert — an alert fires at 3am and the on-call engineer has no idea what to do. The alert email says 'error rate high' with no link to a runbook, dashboard, or mitigation steps. Every alert must include a link to a runbook that specifies: (1) what to check, (2) how to diagnose, (3) how to mitigate, (4) who to escalate to.

## In Production

Alerting mindset is the foundation of SRE practice. Google's SRE book defines the standard: every alert must be: (1) urgent (requires immediate action), (2) actionable (the responder knows what to do), and (3) effective (the action mitigates the problem). Major organizations (Google, Netflix, GitHub) use burn-rate alerting based on SLOs. The 'you are not Google' principle: small teams should have <10 critical alerts; each alert should fire <1x per quarter.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-12-11`.
