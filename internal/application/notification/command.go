package notification

type SendNotificationCommand struct {
	Token string
	Title string
	Body  string
}
