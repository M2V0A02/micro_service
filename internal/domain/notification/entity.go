package notification

import (
	"errors"
	"time"
)

type Notification struct {
	ID     string
	Token  string
	Title  string
	Body   string
	Status Status
	SentAt time.Time
}

type Status string

const (
	StatusPending Status = "pending"
	StatusSent    Status = "sent"
	StatusFailed  Status = "failed"
)

func (n *Notification) Validate() error {
	if n.Token == "" || n.Title == "" || n.Body == "" {
		return errors.New("token, title, and body are required")
	}

	if n.SentAt.IsZero() {
		return errors.New("sent at date is required")
	}

	return nil
}
