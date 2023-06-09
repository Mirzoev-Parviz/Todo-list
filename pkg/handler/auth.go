package handler

import (
	"errors"
	"net/http"
	"test/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input model.User

	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	// Hashing password
	userId, err := h.services.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": userId})
}

type singInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input singInput

	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	// Hashing password
	token, err := h.services.GenerateToken(input.Login, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("error while parsing token")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("error while parsing token")
	}
	return idInt, nil
}
