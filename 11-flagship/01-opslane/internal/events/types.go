// Copyright (c) 2026 Rasel Hossen
// See LICENSE for usage terms.

package events

import (
	"time"

	"github.com/rasel9t6/the-go-engineer/11-flagship/01-opslane/internal/models"
)

type Type string

const (
	TypeOrderCreated          Type = "order.created"
	TypeOrderStatusChanged    Type = "order.status_changed"
	TypePaymentRequested      Type = "payment.requested"
	TypePaymentSettled        Type = "payment.settled"
	TypePaymentFailed         Type = "payment.failed"
	TypeNotificationRequested Type = "notification.requested"
)

// Event is the small, explicit message that crosses the async boundary.
type Event struct {
	ID                string
	Type              Type
	TenantID          int64
	UserID            int64
	OrderID           int64
	PaymentID         int64
	OrderStatus       models.OrderStatus
	PaymentStatus     models.PaymentStatus
	ProviderReference string
	AmountCents       int64
	OccurredAt        time.Time
	Metadata          map[string]string
}

func (e Event) WithOccurredAt(now time.Time) Event {
	if e.OccurredAt.IsZero() {
		e.OccurredAt = now
	}

	return e
}
