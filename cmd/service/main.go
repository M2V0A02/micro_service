package main

import (
	"notification/cmd"
)

func main() {
	container := cmd.NewInternal(cmd.NewContainer())
	// migrator := container.GetMigrator()
	// err := migrator.MigratorUp()
	globalCtx := container.GetGlobalContext()
	server := container.GetServer()
	log := container.GetLogger()

	ctxFields := map[string]string{
		"path": "cmd/service/main.go",
		"name": "main",
	}

	ctx := log.WithFields(globalCtx, ctxFields)
	server.Run(ctx)
}
