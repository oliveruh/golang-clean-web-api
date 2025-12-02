package api

import (
	"fmt"

	"golang-clean-web-api/api/middleware"
	"golang-clean-web-api/api/router"
	"golang-clean-web-api/config"
	"golang-clean-web-api/pkg/logging"

	"github.com/gin-gonic/gin"
)

var logger = logging.NewLogger(config.GetConfig())

func InitServer(cfg *config.Config) {
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors(cfg))

	RegisterRoutes(r, cfg)

	logger := logging.NewLogger(cfg)
	logger.Info(logging.General, logging.Startup, "Server starting", nil)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		logger.Fatal(logging.General, logging.Startup, err.Error(), nil)
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		// Health check
		health := v1.Group("/health")
		router.Health(health)

		// Basic CRUD endpoints
		countries := v1.Group("/countries")
		cities := v1.Group("/cities")
		colors := v1.Group("/colors")

		router.Country(countries, cfg)
		router.City(cities, cfg)
		router.Color(colors, cfg)
	}
}
