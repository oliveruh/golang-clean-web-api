package router

import (
	"github.com/gin-gonic/gin"
	"github.com/naeemaei/golang-clean-web-api/api/handler"
	"github.com/naeemaei/golang-clean-web-api/config"
)

const GetByFilterExp string = "/get-by-filter"

func Country(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCountryHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST(GetByFilterExp, h.GetByFilter)
}

func City(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewCityHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST(GetByFilterExp, h.GetByFilter)
}



func Color(r *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewColorHandler(cfg)

	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
	r.GET("/:id", h.GetById)
	r.POST(GetByFilterExp, h.GetByFilter)
}


