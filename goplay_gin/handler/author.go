package handler

import (
	"fmt"
	"net/http"

	"github.com/404th/goplay_gin/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Author godoc
// @ID create-author
// @Router /author [POST]
// @Summary create author
// @Description create author
// @Tags author
// @Accept json
// @Produce json
// @Param CreateAuthor body model.CreateAuthor true "Create Author"
// @Success 200 {object} model.Response "OK"
// @Response 400 {object} model.Response "Bad Request"
func (h Handler) CreateAuthor(ctx *gin.Context) {
	var (
		author model.Author
	)

	if err := ctx.ShouldBindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": model.Response{
				Data:    err.Error(),
				Message: "cannot bind json",
			},
		})
		return
	}

	id := uuid.NewString()
	author.ID = id

	resp, err := h.strg.Author().CreateAuthor(author.ID, author.Firstname, author.Secondname)
	if err != nil {
		ctx.JSON(http.StatusCreated, gin.H{
			"response": model.Response{
				Data:    fmt.Errorf("cannot create author: %w", err),
				Message: "author not created",
			},
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"response": model.Response{
			Data:    resp,
			Message: "author created",
		},
	})
}

// Author godoc
// @ID get-an-author
// @Router /author/{id} [GET]
// @Summary get an author
// @Description get an author
// @Tags author
// @Accept json
// @Produce json
// @Param id path string true "Get an author"
// @Success 200 {object} model.Response "OK"
// @Response 400 {object} model.Response "Bad Request"
// @Response 404 {object} model.Response "Not found"
func (h Handler) GetAuthorByID(ctx *gin.Context) {
	get_id := ctx.Param("id")

	resp, err := h.strg.Author().GetAuthor(get_id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"response": model.Response{
				Data:    err.Error(),
				Message: "error while getting an author",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": model.Response{
			Data:    resp,
			Message: "author found",
		},
	})
}

func (h Handler) GetAllAuthor(ctx *gin.Context) {
	offset, offset_exists := ctx.GetQuery("offset")
	if !offset_exists {
		offset = "10"
	}

	limit, limit_exists := ctx.GetQuery("limit")
	if !limit_exists {
		limit = "10"
	}

	search, search_exists := ctx.GetQuery("search")
	if !search_exists {
		limit = "10"
	}

	resp, err := h.strg.Author().GetAllAuthor(offset, limit, search)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"response": model.Response{
				Data:    fmt.Errorf("error occured while getting all authors: %w", err),
				Message: "could not get all authors",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": model.Response{
			Data:    resp,
			Message: "got all author",
		},
	})
}

func (h Handler) UpdateAuthor(ctx *gin.Context) {
	var updAuthor model.UpdateAuthor
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&updAuthor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": model.Response{
				Data:    fmt.Errorf("cannot bind json: %w", err),
				Message: "could not bind json while updating author",
			},
		})
		return
	}

	err := h.strg.Author().UpdateAuthor(id, updAuthor)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"response": model.Response{
				Data:    fmt.Errorf("error while updating author: %w", err),
				Message: "could not update author",
			},
		})
		return
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"response": model.Response{
			Data:    nil,
			Message: "author updated",
		},
	})
}

func (h Handler) DeleteAuthor(ctx *gin.Context) {
	id := ctx.Param("id")

	err := h.strg.Author().DeleteAuthor(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"response": model.Response{
				Data:    fmt.Errorf("error while deleting an author: %w", err),
				Message: "could not delete author",
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": model.Response{
			Data:    nil,
			Message: "author deleted successfully",
		},
	})
}
