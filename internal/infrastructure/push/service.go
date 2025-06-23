package push

import "context"

type Service interface {
	SendPush(ctx context.Context, token, title, body string) error
}
