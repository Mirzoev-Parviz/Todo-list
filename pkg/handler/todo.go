package handler

import (
	"net/http"
	"strconv"
	"test/model"

	"github.com/gin-gonic/gin"
)

type taskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        string `json:"done"`
}

func (h *Handler) CreateTODO(c *gin.Context) {
	var todo model.TODO

	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateTODO(userId, todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": id})
}

func (h *Handler) GetTODOS(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	todoList, err := h.services.GetTODOS(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todoList)
}

func (h *Handler) UpdateTODO(c *gin.Context) {
	id, err := getTaskId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var todo model.TODO
	if err = c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err = h.services.UpdateTODO(todo, id, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated successfullt"})
}

func (h *Handler) DeleteTODO(c *gin.Context) {
	id, err := getTaskId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err = h.services.DeleteTODO(id, userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}

func getTaskId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}

	return id, nil
}
