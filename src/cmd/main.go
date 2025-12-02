package main

import (
	"golang-clean-web-api/api"
	"golang-clean-web-api/config"
	"golang-clean-web-api/infra/cache"
	database "golang-clean-web-api/infra/persistence/database"
	"golang-clean-web-api/infra/persistence/migration"
	"golang-clean-web-api/pkg/logging"
)

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
