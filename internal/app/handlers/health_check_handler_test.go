package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler_HealthCheck(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	handler := NewHealthCheckHandler()
	router.GET("/health", handler.HealthCheck())

	// Create a test request
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response body
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Assert response content
	assert.Equal(t, "I'm alive", response["message"])
}

func TestHealthCheckHandler_Name(t *testing.T) {
	handler := NewHealthCheckHandler()
	assert.Equal(t, "health-check-handler", handler.Name())
}
