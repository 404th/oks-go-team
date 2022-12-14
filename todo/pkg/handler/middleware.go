package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	splittedHeader := strings.Split(header, " ")
	if len(splittedHeader) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid token")
		return
	}

	userId, err := h.service.Authorization.ParseToken(splittedHeader[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	userId, exists := c.Get(userCtx)
	if !exists {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	id, ok := userId.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "invalid type of user id")
		return 0, errors.New("invalid user id")
	}

	return id, nil
}
