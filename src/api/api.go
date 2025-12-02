package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/api/middleware"
	"github.com/naeemaei/golang-clean-web-api/config"
)

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors(cfg))

	RegisterRoutes(r)

	log.Printf("Server starting on port %s", cfg.Server.Port)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.GET("/health", Health)
	}
}
