# Channel ownership

## Mission

Understand and apply Channel ownership in the context of professional Go software engineering.

## Prerequisites

- core-11-15

## Mental Model

Every channel has exactly one owner: the goroutine that created it. The owner sends values to the channel and closes it when done. All other goroutines are receivers: they read from the channel until it is closed. The ownership contract is: the owner creates, sends, and closes; receivers only receive. This prevents send-on-closed and double-close panics, and ensures receivers can terminate their range loops cleanly. Channels are not shared state — they are communication pipelines, and pipelines have an operator (the owner) who controls the flow.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

A channel in Go is a runtime object (hchan struct) allocated on the heap. The hchan contains a mutex, a circular buffer (for buffered channels), send/recv wait queues (lists of goroutines blocked on the channel), and a closed flag. When close(ch) is called, the runtime sets the closed flag, dequeues all goroutines from the recvq (receive wait queue), and schedules them to run — each receives the zero value with ok=false. Any goroutine in the sendq (send wait queue) receives a 'panic: send on closed channel' when it runs. The channel is garbage collected when no goroutine holds a reference to it and all blocked goroutines have been woken. The ownership discipline is not enforced by the runtime — any goroutine with a reference can close the channel — but violating it causes panics and races. The discipline is enforced by convention, documentation, and code review.

## Run Instructions

```bash
go run ./curriculum/modules/11-lifecycle-context-concurrency/lessons/16-channel-ownership
go test ./curriculum/modules/11-lifecycle-context-concurrency/lessons/16-channel-ownership
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Allowing multiple goroutines to send to or close the same channel without coordination — a channel must have exactly one owner (the goroutine that creates and closes it); multiple senders without synchronization cause panics (send on closed channel) or data races.
- Closing a channel from the receiver side — the receiver does not know if all senders are done; closing from the receiver causes the sender to panic on the next send. The owner (sender) should always close the channel.
- Sending on a closed channel panics — but checking ch != nil before sending does not help; the channel must be open. The correct pattern is: the owner signals completion by closing, and receivers use the comma-ok idiom or range.
- Not closing a channel at all — the receiver goroutines range over the channel and never exit, causing a goroutine leak. The owner must close the channel when all sends are done.
- Sharing channel ownership across package boundaries — channels should be created in one place and the ownership (who closes) should be documented or enforced by the API; unexported channels with public send/receive functions enforce ownership.

## In Production

Every production Go system that uses channels follows the ownership pattern. Database connection pools: the pool owner creates a channel of connections and closes it on shutdown. Pipeline stages: each stage owns its output channel and closes it when processing is complete. gRPC streams: the server owns the send channel and closes it on completion or error. Kubernetes controllers: informer event channels are owned by the informer and closed on shutdown. The ownership pattern is baked into Go's concurrency idioms.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-11-17`.
