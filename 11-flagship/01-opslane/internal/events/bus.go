// Copyright (c) 2026 Rasel Hossen
// See LICENSE for usage terms.

package events

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	ErrInvalidEvent = errors.New("invalid event")
	ErrQueueFull    = errors.New("event queue full")
	ErrBusClosed    = errors.New("event bus closed")
)

type Bus struct {
	events     chan Event
	closed     chan struct{}
	once       sync.Once
	now        func() time.Time
	mu         sync.RWMutex
	closedFlag bool
}

func NewBus(capacity int) *Bus {
	if capacity <= 0 {
		capacity = 1
	}

	return &Bus{
		events: make(chan Event, capacity),
		closed: make(chan struct{}),
		now:    time.Now,
	}
}

func (b *Bus) Subscribe() <-chan Event {
	if b == nil {
		return nil
	}

	return b.events
}

func (b *Bus) TryPublish(event Event) error {
	if b == nil {
		return fmt.Errorf("event bus is not configured")
	}

	// Fast-path: if the bus is already closed, reject immediately
	b.mu.RLock()
	if b.closedFlag {
		b.mu.RUnlock()
		return ErrBusClosed
	}
	b.mu.RUnlock()

	prepared, err := b.prepare(event)
	if err != nil {
		return err
	}

	select {
	case b.events <- prepared:
		return nil
	case <-b.closed:
		return ErrBusClosed
	default:
		return ErrQueueFull
	}
}

func (b *Bus) Publish(ctx context.Context, event Event) error {
	if b == nil {
		return fmt.Errorf("event bus is not configured")
	}
	if ctx == nil {
		ctx = context.Background()
	}

	// Fast-path: reject if bus is already closed
	b.mu.RLock()
	if b.closedFlag {
		b.mu.RUnlock()
		return ErrBusClosed
	}
	b.mu.RUnlock()

	prepared, err := b.prepare(event)
	if err != nil {
		return err
	}

	// Check if bus is closed first to avoid race with Close
	select {
	case <-b.closed:
		return ErrBusClosed
	default:
		// Proceed to try sending with context
	}

	// Try to send the event with context
	select {
	case b.events <- prepared:
		return nil
	case <-b.closed:
		return ErrBusClosed
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (b *Bus) Close() {
	if b == nil {
		return
	}
	b.once.Do(func() {
		// Mark closed in a thread-safe way and signal close to waiters
		b.mu.Lock()
		if !b.closedFlag {
			b.closedFlag = true
			close(b.closed)
		}
		b.mu.Unlock()
	})
}

func (b *Bus) prepare(event Event) (Event, error) {
	if event.Type == "" || event.TenantID <= 0 {
		return Event{}, ErrInvalidEvent
	}

	return event.WithOccurredAt(b.now().UTC()), nil
}
