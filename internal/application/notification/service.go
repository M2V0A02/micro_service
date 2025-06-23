package notification

import (
	"context"
	"notification/internal/domain/notification"
	"time"
)

type Service struct {
	domainService *notification.Service
}

func NewService(domainService *notification.Service) *Service {
	return &Service{domainService: domainService}
}

func (s *Service) SendNotification(ctx context.Context, cmd SendNotificationCommand) error {
	aggregate := notification.NewNotificationAggregate(cmd.Token, cmd.Title, cmd.Body, time.Now())
	return s.domainService.SendNotification(ctx, aggregate)
}
