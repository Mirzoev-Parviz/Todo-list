package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user_Id"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, "empty auth header")
		return
	}

	headerPars := strings.Split(header, " ")

	if len(headerPars) != 2 {
		c.JSON(http.StatusUnauthorized, "Invalid auth header")
		return
	}

	// fmt.Println("headerPars[1]", headerPars[1])
	userID, err := h.services.Authorization.ParseToken(headerPars[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userID)
}
