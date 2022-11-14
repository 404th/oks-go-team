package handler

import (
	"github.com/404th/goplay_gin/config"
	"github.com/404th/goplay_gin/storage"
)

type Handler struct {
	strg storage.StorageI
	cfg  config.Config
}

func NewHandler(strg storage.StorageI, cfg config.Config) *Handler {
	return &Handler{
		strg: strg,
		cfg:  cfg,
	}
}

// func (h *Handler) parseLimitQueryParam(c *gin.Context) (int, error) {
// 	return strconv.Atoi(c.DefaultQuery("limit", h.cfg.DEFAULTLIMIT))
// }

// func (h *Handler) parseOffsetQueryParam(c *gin.Context) (int, error) {
// 	return strconv.Atoi(c.DefaultQuery("offset", h.cfg.DEFAULTOFFSET))
// }
