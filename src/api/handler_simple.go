package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /v1/health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Server is running",
	})
}
