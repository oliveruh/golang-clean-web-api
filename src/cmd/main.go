package main

import (
	"github.com/naeemaei/golang-clean-web-api/api"
	"github.com/naeemaei/golang-clean-web-api/config"
)

func main() {
	cfg := config.GetConfig()
	api.InitServer(cfg)
}
