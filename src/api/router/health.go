package router

import (
	"golang-clean-web-api/api/handler"

	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handler.NewHealthHandler()

	r.GET("/", handler.Health)
}
