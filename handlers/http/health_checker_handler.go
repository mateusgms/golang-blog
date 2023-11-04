package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthHandler is a handler for health check
func HealthHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}