package handler

import (
	"net/http"

	"github.com/404th/todo/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

}

func (h *Handler) signIn(c *gin.Context) {}
