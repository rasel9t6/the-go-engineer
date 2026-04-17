# 10 Production

## Mission

This stage teaches how Go services behave once they are running for real: logging, shutdown, and
deployment shape the program as much as the business logic does.

By the end of this stage, a learner should be able to:

- reason about operational log shape and request context
- shut services down without losing control of in-flight work
- understand how packaging and deployment choices affect runtime behavior

## Stage Map

| Track | Surface | Core Job |
| --- | --- | --- |
| `PRD.1` | [structured logging](./01-structured-logging) | teach operational log shape and redaction |
| `PRD.2` | [graceful shutdown](./02-graceful-shutdown) | teach lifecycle safety and drain order |
| `PRD.3` | [docker and deployment](./03-docker-and-deployment) | teach packaging and deployment-oriented workflows |

## Suggested Learning Flow

1. Start with [structured logging](./01-structured-logging).
2. Continue into [graceful shutdown](./02-graceful-shutdown).
3. Use [docker and deployment](./03-docker-and-deployment) to connect the code to delivery concerns.

## Next Step

After this stage, move to [11 Flagship](../11-flagship).
