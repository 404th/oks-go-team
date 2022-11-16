package handler

import (
	"fmt"
	"net/http"

	"github.com/404th/todo/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var todolist model.TodoList
	if err = c.ShouldBindJSON(&todolist); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("++++++++++++++++")
	fmt.Println(userId)
	fmt.Println("++++++++++++++++")

	fmt.Println("++++++++++++++++")
	fmt.Println(todolist)
	fmt.Println("++++++++++++++++")

	todolist_id, err := h.service.CreateList(userId, todolist)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": todolist_id,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {}

func (h *Handler) getListById(c *gin.Context) {}

func (h *Handler) updateList(c *gin.Context) {}

func (h *Handler) deleteList(c *gin.Context) {}
