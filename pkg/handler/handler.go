package handler

import (
	"congo/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	accounts := router.Group("/accounts")
	{
		accounts.GET("/", h.getAll)
		//accounts.GET("filter", h.filter)
		filter := accounts.Group("/filter")
		{
			filter.GET("/", h.filter)

			// filter.GET(":email")
			// filter.GET(":status")
			// filter.GET(":fname")
			// filter.GET(":sname")
			// filter.GET(":phone")
			// filter.GET(":country")
			// filter.GET(":city")
			// filter.GET(":birth")
			// filter.GET(":sex")
			// filter.GET(":likes")
			// filter.GET(":premium")
		}
	}

	return router
}
