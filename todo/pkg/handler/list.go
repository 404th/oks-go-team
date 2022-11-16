package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	userId, exists := c.Get(userCtx)
	if !exists {
		newErrorResponse(c, http.StatusUnauthorized, "unauthorized user")
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"id": userId,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {}

func (h *Handler) getListById(c *gin.Context) {}

func (h *Handler) updateList(c *gin.Context) {}

func (h *Handler) deleteList(c *gin.Context) {}
