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

func (s *Service) SendScheduleNotification(ctx context.Context, cmd SendScheduleNotificationCommand) error {
	aggregate := notification.NewNotificationAggregate(cmd.Token, cmd.Title, cmd.Body, cmd.SentAt)
	return s.domainService.SendScheduleNotification(ctx, aggregate)
}
