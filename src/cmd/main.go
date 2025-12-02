package main

import (
	"golang-clean-web-api/api"
	"golang-clean-web-api/config"
	"golang-clean-web-api/infra/cache"
	database "golang-clean-web-api/infra/persistence/database"
	"golang-clean-web-api/infra/persistence/migration"
	"golang-clean-web-api/pkg/logging"

	_ "golang-clean-web-api/docs" // This line is necessary for Swagger to find your docs
)

// @title Golang Clean Web API
// @version 1.0
// @description A clean architecture web API with JWT authentication, rate limiting, and comprehensive documentation
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}

	err = database.InitDb(cfg)
	defer database.CloseDb()
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	migration.Up1()

	api.InitServer(cfg)
}
