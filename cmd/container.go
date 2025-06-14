package cmd

import (
	"context"
	"log"

	"notification/migrations"
	"notification/pkg/logger"

	embPg "github.com/fergusstrange/embedded-postgres"
	"github.com/jmoiron/sqlx"
)

// контейнер внешних зависимостей приложения
// тут мы инициализируем все инфраструктурные зависимости
type Container struct {
	gCtx             context.Context
	configuration    *configuration
	db               *sqlx.DB
	embeddedPostgres *embPg.EmbeddedPostgres
	migrator         *migrations.Migrator
	logger           *logger.Logger
}

func NewContainer() *Container {
	return &Container{
		configuration: newFromEnv(),
	}
}

// для доступа внутренних зависимостей к конфигурации
func (e *Container) GetConfiguration() *configuration {
	return e.configuration
}

func (e *Container) GetGlobalContext() context.Context {
	if e.gCtx == nil {
		e.gCtx = context.Background()
	}

	return e.gCtx
}

func (e *Container) GetPostgres() *sqlx.DB {
	if e.db == nil {
		var err error
		e.db, err = NewSqlxConn(e.configuration.GetPostgresConfiguration())
		if err != nil {
			log.Fatal(err)
		}
	}

	return e.db
}

func (e *Container) GetEmbeddedPostgres() *embPg.EmbeddedPostgres {
	if e.embeddedPostgres == nil {
		e.embeddedPostgres = embPg.NewDatabase(
			e.configuration.
				GetPostgresConfiguration().
				GetEmbeddedPostgresConfig(),
		)
	}

	return e.embeddedPostgres
}

func (e *Container) GetLogger() *logger.Logger {
	if e.logger == nil {
		e.logger = logger.New()
	}

	return e.logger
}

func (e *Container) GetMigrator() *migrations.Migrator {
	if e.migrator == nil {
		e.migrator = migrations.NewMigrator(
			e.configuration.
				GetPostgresConfiguration().
				GetMigrateConnectionString(),
		)
	}

	return e.migrator
}
