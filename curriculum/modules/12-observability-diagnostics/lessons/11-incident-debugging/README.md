# Incident debugging

## Mission

Understand and apply Incident debugging in the context of professional Go software engineering.

## Prerequisites

- core-12-10

## Mental Model

Incident debugging is a data-driven investigation, not a guessing game. The process is: observe (check dashboards and alerts) → correlate (find correlation IDs and trace spans) → isolate (identify the failing component or layer) → diagnose (find the root cause using logs, profiles, and dumps) → remediate (fix, roll back, or scale). Each step uses different observability data — never jump to diagnosis without first observing the blast radius.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

Go's runtime provides pprof profiles at /debug/pprof/: goroutine (all goroutine stacks), heap (memory allocation), profile (CPU), mutex (contended mutexes), and block (blocking operations). When debugging an incident, the engineer fetches a goroutine profile to see if goroutines are stuck in Gwaiting or Gsyscall — a thousand goroutines waiting on a channel receive indicates a blocked producer. The mutex profile shows which mutexes are contended and which call sites hold them. The CPU profile shows which function is consuming CPU — useful for infinite loops or busy-waiting.

## Run Instructions

```bash
go run ./curriculum/modules/12-observability-diagnostics/lessons/11-incident-debugging
go test ./curriculum/modules/12-observability-diagnostics/lessons/11-incident-debugging
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Starting an incident investigation by guessing the root cause rather than gathering data — guessing leads to confirmation bias, wasted time on wrong subsystems, and delayed resolution.
- Ignoring the alert dashboard and jumping straight into code — the dashboard shows the blast radius, affected users, and timeline; skipping it means investigating blind.
- Focusing on a single log line instead of the full correlation-ID-scoped trace — one log line out of context is noise; the full trace shows the sequence of events.
- Making changes to production (restarting, scaling) before understanding the root cause — the change may mask the symptom temporarily but does not prevent recurrence.
- Not communicating status updates during the incident — stakeholders, teammates, and affected users are left guessing, eroding trust and causing parallel investigation.

## In Production

Every production incident at every Go shop follows the same pattern. Stripe's incident response process, Google's SRE handbook, and the Go community's postmortem culture all emphasize: gather data before acting, use correlation IDs to trace across services, and write postmortems without blame. The difference between a 10-minute and a 2-hour incident is almost always whether the engineer looked at the dashboard first or started guessing.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-12-12`.
