package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"project/internal/auth"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ah *APIHandler) LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := ah.userService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
