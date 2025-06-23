package main

import (
	"context"
	"notification/pkg/di"
)

func main() {
	container := di.NewContainer()
	ctx := container.GetLogger().WithFields(context.Background(), map[string]string{
		"path": "cmd/service/main.go",
		"name": "main",
	})
	server := container.GetServer()
	server.Run(ctx)
}
