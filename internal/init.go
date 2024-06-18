package internal

import (
	"context"
	"fmt"

	_ "github.com/bigxxby/get-counter-test/docs"
	"github.com/bigxxby/get-counter-test/internal/handlers"
	"github.com/bigxxby/get-counter-test/internal/repository"
	"github.com/bigxxby/get-counter-test/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run(db *redis.Client, router *gin.Engine) {
	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handlers.NewHandler(service)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	api.Use(gin.Logger())
	api.GET("/counter", handler.CounterHandler)
	api.DELETE("/clear", handler.ClearCounterHandler)

	fmt.Println("Server is running on port 8080")
	router.Run()
}

func Connection() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	fmt.Println("Connected to Redis")

	return client, nil
}
