package repository

import (
	"context"
	"notification/internal/domain/notification"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type NotificationRepository struct {
	db *sqlx.DB
	qb squirrel.StatementBuilderType
}

func NewNotificationRepository(db *sqlx.DB) *NotificationRepository {
	return &NotificationRepository{
		db: db,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *NotificationRepository) Save(ctx context.Context, notification *notification.Notification) error {
	query := r.qb.Insert("notifications").
		Columns("id", "token", "title", "body", "status").
		Values(notification.ID, notification.Token, notification.Title, notification.Body, notification.Status).
		Suffix("ON CONFLICT (id) UPDATE SET token = EXCLUDED.token, title = EXCLUDED.title, body = EXCLUDED.body, status = EXCLUDED.status")

	sql, args, _ := query.ToSql()
	_, err := r.db.ExecContext(ctx, sql, args...)
	return err
}

func (r *NotificationRepository) GetByID(ctx context.Context, id string) (*notification.Notification, error) {
	var n notification.Notification
	query := r.qb.Select("id", "token", "title", "body", "status").From("notifications").Where(squirrel.Eq{"id": id})
	sql, args, _ := query.ToSql()
	err := r.db.GetContext(ctx, &n, sql, args...)
	return &n, err
}
