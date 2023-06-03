package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"project/internal/users"
)

type AuthService struct {
	userRepo users.Repository
}

func NewAuthService(userRepo users.Repository) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (ah *AuthService) Login(c *gin.Context) {
	var loginData LoginRequest
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "(Service)Invalid request payload"})
		return
	}

	// Retrieve the user by username from the repository
	user, err := ah.userRepo.GetUserByUsernameAndPassword(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare the provided password with the hashed password from the user model
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	tokenString, err := GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the token as the authentication result
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
