package main

import (
	"context"
	"log"
	"notification/pkg/di"

	"github.com/joho/godotenv"
)

func init() {
	// Загрузка переменных из .env в окружение os.Getenv
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Не удалось загрузить .env:", err)
	} else {
		log.Println(".env загружен")
	}
}

func main() {
	container := di.NewContainer()
	ctx := container.GetLogger().WithFields(context.Background(), map[string]string{
		"path": "cmd/service/main.go",
		"name": "main",
	})
	server := container.GetServer()
	server.Run(ctx)
}
