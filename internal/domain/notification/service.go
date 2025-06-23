package notification

import (
	"context"
	"notification/internal/infrastructure/push"
)

type Service struct {
	repo        Repository
	pushService push.Service
}

func NewService(repo Repository, pushService push.Service) *Service {
	return &Service{repo: repo, pushService: pushService}
}

func (s *Service) SendNotification(ctx context.Context, notification *NotificationAggregate) error {
	if err := notification.Validate(); err != nil {
		return err
	}

	if err := s.repo.Save(ctx, &notification.Notification); err != nil {
		return err
	}

	if err := s.pushService.SendPush(ctx, notification.Token, notification.Title, notification.Body); err != nil {
		notification.MarkAsFailed()
		_ = s.repo.Save(ctx, &notification.Notification)
		return err
	}

	notification.MarkAsSent()
	return s.repo.Save(ctx, &notification.Notification)
}

func (s *Service) SendScheduleNotification(ctx context.Context, notification *NotificationAggregate) error {
	if err := notification.Validate(); err != nil {
		return err
	}

	if err := s.repo.Save(ctx, &notification.Notification); err != nil {
		return err
	}

	return nil
}
