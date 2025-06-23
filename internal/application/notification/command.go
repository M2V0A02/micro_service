package notification

import "time"

type SendNotificationCommand struct {
	Token string
	Title string
	Body  string
}

type SendScheduleNotificationCommand struct {
	Token  string
	Title  string
	Body   string
	SentAt time.Time
}
