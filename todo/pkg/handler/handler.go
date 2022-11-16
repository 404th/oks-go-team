package handler

import (
	"github.com/404th/todo/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	// auth
	r.POST("/auth/sign-up", h.signUp)
	r.GET("/auth/sign-in", h.signIn)

	api := r.Group("/api", h.userIdentity)
	{
		// lists
		api.GET("/lists", h.getAllLists)
		api.POST("/lists", h.createList)
		api.GET("/lists/:id", h.getListById)
		api.PUT("/lists/:id", h.updateList)
		api.DELETE("/lists/:id", h.deleteList)

		// items
		api.GET("/lists/:id/items", h.getAllItems)
		api.POST("/lists/:id/items", h.createItem)
		api.GET("/lists/:id/items/:item_id", h.getItemById)
		api.PUT("/lists/:id/items/:item_id", h.updateItem)
		api.DELETE("/lists/:id/items/:item_id", h.deleteItem)
	}

	return r
}
