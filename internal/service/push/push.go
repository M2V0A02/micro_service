package push

import (
	"context"
	"fmt"
	"notification/internal/deps"
)

type MockPushService struct {
	logger deps.Logger
}

func NewMockPushService(logger deps.Logger) *MockPushService {
	return &MockPushService{
		logger: logger,
	}
}

func (s *MockPushService) SendPush(ctx context.Context, token, title, body string) error {
	msg := fmt.Sprintf("Mock push sent: token=%s, title=%s, body=%s", token, title, body)
	fmt.Println(msg)
	s.logger.Info(ctx, msg)
	return nil
}
