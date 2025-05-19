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

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Returns a status message indicating the service is running
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Service is running"
// @Router /api/health-check [get]
func (h *HealthCheckHandler) HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm alive",
		})
	}
}
