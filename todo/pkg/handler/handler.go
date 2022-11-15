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
	r.POST("/auth/sign-in", h.signIn)

	// lists
	r.GET("/api/lists", h.getAllLists)
	r.POST("/api/lists", h.createList)
	r.GET("/api/lists/:id", h.getListById)
	r.PUT("/api/lists/:id", h.updateList)
	r.DELETE("/api/lists/:id", h.deleteList)

	// items
	r.GET("/api/lists/:id/items", h.getAllItems)
	r.POST("/api/lists/:id/items", h.createItem)
	r.GET("/api/lists/:id/items/:item_id", h.getItemById)
	r.PUT("/api/lists/:id/items/:item_id", h.updateItem)
	r.DELETE("/api/lists/:id/items/:item_id", h.deleteItem)

	return r
}
