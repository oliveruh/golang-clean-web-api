package router

import (
	"golang-clean-web-api/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Swagger(r *gin.Engine, cfg *config.Config) {
	// Only enable Swagger in debug/development mode
	if cfg.Server.RunMode == "debug" || cfg.Server.RunMode == "development" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
