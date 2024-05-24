package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/YasserRABIE/authentication-porject/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func HandleAuth(c *gin.Context) {
	userName, _ := c.Get("userName")
	fmt.Println("username:", userName)

	token, err := generateJWTToken(userName)
	if err != nil {
		resBody := models.NewFailedResponse(400, map[string]string{
			"error": "Failed to generate token",
		})

		c.JSON(http.StatusBadRequest, resBody)
		return
	}

	resBody := models.NewSuccessResponse(200, map[string]interface{}{
		"token": token,
	})

	c.JSON(http.StatusOK, &resBody)
}

func generateJWTToken(username interface{}) (string, error) {
	claims := jwt.MapClaims{
		"userName": username,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(), // Token expiration time (30 days from now)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
