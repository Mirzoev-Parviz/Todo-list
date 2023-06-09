package handler

import (
	"test/pkg/service"

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
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up/", h.signUp)
		auth.POST("/sign-in", h.signIn)

	}
	api := router.Group("/api", h.userIdentity)
	{

		todo := api.Group("/todo")
		{
			todo.POST("/", h.CreateTODO)
			todo.GET("/", h.GetTODOS)
			todo.PUT("/:id", h.UpdateTODO)
			todo.DELETE("/:id", h.DeleteTODO)
		}
	}

	return router
}
