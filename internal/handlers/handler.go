package handlers

import (
	"log"

	"github.com/bigxxby/get-counter-test/internal/service"
	"github.com/gin-gonic/gin"
)

// Handler описывает интерфейс обработчиков.
type Handler interface {
	CounterHandler(c *gin.Context)
	ClearCounterHandler(c *gin.Context)
}

type Message struct {
	Message string `json:"message"`
}
type handler struct {
	service service.Service
}

// NewHandler создает новый экземпляр обработчика с заданной сервисной службой.
func NewHandler(service service.Service) *handler {
	return &handler{
		service: service,
	}
}

// @Summary Получить текущее значение счетчика
// @Description Получение текущего значения счетчика из сервиса
// @Tags Counter
// @Accept json
// @Produce json
// @Success 200 {object} Message
// @Failure 500 {object} Message
// @Router /api/counter [get]
func (h *handler) CounterHandler(c *gin.Context) {
	count, err := h.service.Counter()
	if err != nil {
		log.Printf("Failed to get count: %v", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"count": count})
}

// @Summary Сбросить значение счетчика
// @Description Сброс текущего значения счетчика в сервисе
// @Tags Counter
// @Accept json
// @Produce json
// @Success 200 {object} Message
// @Failure 500 {object} Message
// @Router /api/clear [delete]
func (h *handler) ClearCounterHandler(c *gin.Context) {
	err := h.service.ClearCounter()
	if err != nil {
		log.Printf("Failed to delete count: %v", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(200, gin.H{"message": "Counter deleted"})
}
