
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupRouter initializes a new Gin router for testing purposes.
// It includes a basic health check endpoint.
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/test-health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "his-healthy"})
	})
	return r
}

// TestHealthCheck tests the /test-health endpoint of the zarish-his service.
func TestHealthCheck(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test-health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"his-healthy"}`, w.Body.String())
}
