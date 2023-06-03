package auth

import (
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = os.Getenv("SECRET_KEY")

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (string, error) {
	trim_token := strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(trim_token, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	return username, nil
}
