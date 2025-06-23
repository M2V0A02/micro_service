package notification

import (
	"time"

	"github.com/google/uuid"
)

type NotificationAggregate struct {
	Notification
}

func NewNotificationAggregate(token, title, body string, sentAt time.Time) *NotificationAggregate {
	return &NotificationAggregate{
		Notification: Notification{
			ID:     uuid.New().String(),
			Token:  token,
			Title:  title,
			Body:   body,
			Status: StatusPending,
			SentAt: sentAt,
		},
	}
}

func (a *NotificationAggregate) MarkAsSent() {
	a.Status = StatusSent
}

func (a *NotificationAggregate) MarkAsFailed() {
	a.Status = StatusFailed
}
