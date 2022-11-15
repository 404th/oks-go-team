package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	// auth
	r.POST("/auth/sign-up")
	r.POST("/auth/sign-in")

	// lists
	r.GET("/api/lists")
	r.POST("/api/lists")
	r.GET("/api/lists/:id")
	r.PUT("/api/lists/:id")
	r.DELETE("/api/lists/:id")

	// items
	r.GET("/api/lists/:id/items")
	r.POST("/api/lists/:id/items")
	r.GET("/api/lists/:id/items/:item_id")
	r.PUT("/api/lists/:id/items/:item_id")
	r.DELETE("/api/lists/:id/items/:item_id")

	return r
}
