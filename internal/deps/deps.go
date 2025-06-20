package deps

import (
	"context"
)

type Logger interface {
	Info(ctx context.Context, message string, args ...any)
	Error(ctx context.Context, err error, args ...any)
}

type PushService interface {
	SendPush(ctx context.Context, token, title, body string) error
}
