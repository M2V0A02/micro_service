package di

import (
	"context"
	"log"
	applicationNotification "notification/internal/application/notification"
	domainNotification "notification/internal/domain/notification"
	notificationRepository "notification/internal/domain/notification"
	"notification/internal/infrastructure/db/postgres/repository"
	push "notification/internal/infrastructure/push"
	pushMock "notification/internal/infrastructure/push/mock"
	"notification/internal/infrastructure/server"
	"notification/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type Container struct {
	gCtx                           context.Context
	config                         *configuration
	db                             *sqlx.DB
	logger                         *logger.Logger
	notificationRepository         notificationRepository.Repository
	pushService                    push.Service
	applicationNotificationService *applicationNotification.Service
	domainNotificationService      *domainNotification.Service
	server                         *server.Server
}

func NewContainer() *Container {
	config := newFromEnv()
	return &Container{
		gCtx:   context.Background(),
		config: config,
	}
}

func (c *Container) GetLogger() *logger.Logger {
	if c.logger == nil {
		c.logger = logger.New()
	}
	return c.logger
}

func (c *Container) GetPostgres() *sqlx.DB {
	if c.db == nil {
		var err error
		c.db, err = NewSqlxConn(c.config.GetPostgresConfiguration())
		if err != nil {
			log.Fatal(err)
		}
	}
	return c.db
}

func (c *Container) GetNotificationRepository() notificationRepository.Repository {
	if c.notificationRepository == nil {
		c.notificationRepository = repository.NewNotificationRepository(c.GetPostgres())
	}
	return c.notificationRepository
}

func (c *Container) GetPushService() push.Service {
	if c.pushService == nil {
		c.pushService = pushMock.NewMockPushService(c.GetLogger())
	}
	return c.pushService
}

func (c *Container) GetDomainNotificationService() *domainNotification.Service {
	if c.domainNotificationService == nil {
		c.domainNotificationService = domainNotification.NewService(c.GetNotificationRepository(), c.GetPushService())
	}

	return c.domainNotificationService
}

func (c *Container) GetApplicationNotificationService() *applicationNotification.Service {
	if c.applicationNotificationService == nil {
		c.applicationNotificationService = applicationNotification.NewService(c.GetDomainNotificationService())
	}
	return c.applicationNotificationService
}

func (c *Container) GetServer() *server.Server {
	if c.server == nil {
		c.server = server.NewServer(c.GetApplicationNotificationService(), c.GetLogger())
	}
	return c.server
}
