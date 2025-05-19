package handlers

import (
	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {
}

func (h *HealthCheckHandler) Name() string {
	return "health-check-handler"
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm alive",
		})
	}
}
