package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"project/internal/auth"
	"project/internal/users"
)

func (ah *APIHandler) CreateUserHandler(c *gin.Context) {
	var req users.RegisterUser
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := ah.userService.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_profile": user})
}

func (ah *APIHandler) GetProfileHandler(c *gin.Context) {
	c.Request.Header.Set("Authorization", c.GetHeader("Authorization"))
	auth.JWTMiddleware()(c)

	username, _ := auth.VerifyToken(c.GetHeader("Authorization"))
	user, err := ah.userService.Find(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_profile": user})
}

func (ah *APIHandler) GetOrderHistoryHandler(c *gin.Context) {
	c.Request.Header.Set("Authorization", c.GetHeader("Authorization"))
	auth.JWTMiddleware()(c)

	username, _ := c.Get("username")
	user_id, _ := ah.userService.FindID(username.(string))

	orders, err := ah.orderService.GetOrderHistory(*user_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}
