package handler

import (
	"net/http"

	"github.com/404th/goplay_gin/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": model.Response{
			Data:    "Pong",
			Message: "successfully connected",
		},
	})
}
