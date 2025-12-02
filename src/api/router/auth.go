package router

import (
	"golang-clean-web-api/api/handler"
	"golang-clean-web-api/config"

	"github.com/gin-gonic/gin"
)

func Auth(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewAuthHandler(cfg)

	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.POST("/refresh", h.RefreshToken)
}
