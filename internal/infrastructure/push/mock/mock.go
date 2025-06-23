package push

import (
	"context"
	"fmt"
	"notification/pkg/logger"
)

type MockPushService struct {
	logger *logger.Logger
}

func NewMockPushService(logger *logger.Logger) *MockPushService {
	return &MockPushService{logger: logger}
}

func (s *MockPushService) SendPush(ctx context.Context, token, title, body string) error {
	msg := fmt.Sprintf("Mock push sent: token=%s, title=%s, body=%s", token, title,
		body)
	s.logger.Info(ctx, msg)
	return nil
}
