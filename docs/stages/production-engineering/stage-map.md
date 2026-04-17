# Production Engineering Stage Map

## Stage Goal

This stage teaches learners how to think about runtime visibility, shutdown safety, and deployment
workflow as part of engineering correctness.

## Public Stage Shape

| Surface | Role In Stage |
| --- | --- |
| [structured-logging](../../../10-production/01-structured-logging/) | live beta path for operational log shape and redaction |
| [graceful-shutdown](../../../10-production/02-graceful-shutdown/) | live beta path for drain order and lifecycle safety |
| [docker-and-deployment](../../../10-production/03-docker-and-deployment/) | reference surface for packaging and deployment workflow |

## Live Milestone Backbone

| ID | Surface | Why It Matters |
| --- | --- | --- |
| `SL.5` | PII redactor exercise | proves that log pipelines can enforce operational policy |
| `GS.3` | shutdown capstone | proves that service stop behavior can be coordinated deliberately |

## Recommended Order

1. `SL.1` through `SL.5`
2. `GS.1` through `GS.3`
3. `docker-and-deployment`

## Reference Surfaces That Still Matter

These are important production surfaces, but they are not the current public beta proof path:

- `docker-and-deployment`
