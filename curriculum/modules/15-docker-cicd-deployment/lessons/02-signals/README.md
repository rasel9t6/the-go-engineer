# Signals

## Mission

Understand and apply Signals in the context of professional Go software engineering.

## Prerequisites

- core-15-01

## Mental Model

Signals are OS-level interrupts delivered to a process. A signal handler is like an interrupt service routine (ISR): it must be fast, non-blocking, and reentrant. Go's signal.Notify converts OS signals into Go channel messages, allowing goroutines to handle signals synchronously instead of in an interrupt context.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

signal.Notify registers a channel in a global map of signal->channels. When a signal arrives, the Go runtime's signal handler (installed in sigtramp) looks up the registered channels in the map and sends the signal to each channel in a non-blocking select. If the channel is full (unbuffered or buffer full), the signal is dropped — the send is skipped. signal.NotifyContext creates a cancel context and spawns a goroutine that listens on the signal channel, calling cancel() when a signal arrives. The context's cancel function deregisters the signal handler to prevent goroutine leaks. Under the hood, Go replaces the OS signal handler with its own sigtramp handler that queues signals in a global signal ring buffer and wakes up the runtime's signal goroutine via eventfd (Linux) or mach port (macOS).

## Run Instructions

```bash
go run ./curriculum/modules/15-docker-cicd-deployment/lessons/02-signals
go test ./curriculum/modules/15-docker-cicd-deployment/lessons/02-signals
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using an unbuffered channel for signal.Notify — when a signal arrives and the channel is not ready to receive (main goroutine is busy), the signal is dropped. This causes the process to miss SIGTERM and not shut down. Always use a buffered channel of size 1 for signal.Notify.
- Calling signal.Notify without removing the default handler — signal.Notify adds the channel to the signal's list of recipients but does NOT replace the OS default action. For SIGINT and SIGTERM, the default action is to terminate the process. If the channel is not read in time, the default handler still runs and kills the process. Fix: signal.Reset() or signal.Ignore() for signals you want to handle yourself.
- Forgetting that signal.Notify affects all goroutines — calling signal.Notify in one goroutine registers the channel globally. Any signal sent to the process is delivered to that channel, even if another goroutine expected to handle it. Signal handlers are process-wide, not goroutine-scoped.
- Ignoring SIGPIPE — writing to a broken pipe (e.g., writing to stdout when the reader has exited) sends SIGPIPE to the writing process with default action 'terminate'. Go's os/signal package does not intercept SIGPIPE by default, so a Go program writing to a closed stdout is killed silently. Fix: signal.Notify(sigCh, syscall.SIGPIPE) and handle it or ignore it.

## In Production

Every Go service in Kubernetes relies on signal handling for graceful shutdown. Common patterns: SIGTERM triggers graceful shutdown (HTTP drain, worker drain, DB close), SIGHUP triggers config reload (hot-reload without restart), SIGUSR1 triggers log rotation (reopen log files after rotation), SIGUSR2 triggers debug endpoints (pprof, goroutine dump). Opslane's graceful shutdown handles SIGTERM by draining HTTP connections, NACKing queue messages, flushing metrics, and closing the database pool.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-15-03`.
