package main

import (
	"github.com/bigxxby/get-counter-test/internal"
	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация приложения Gin
	router := gin.Default()

	// Инициализация Redis клиента
	redisClient, err := internal.Connection()
	if err != nil {
		panic(err)
	}

	// Запуск приложения
	internal.Run(redisClient, router)
}
