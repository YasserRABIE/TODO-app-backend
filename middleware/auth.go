package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/YasserRABIE/authentication-porject/initializers"
	"github.com/YasserRABIE/authentication-porject/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func HandleAuth(c *gin.Context) {
	userName, _ := c.Get("userName")

	token, err := generateJWTToken(c, userName)
	if err != nil {
		resBody := models.NewFailedResponse(http.StatusUnauthorized, map[string]string{
			"error": "Failed to generate token",
		})

		c.JSON(http.StatusBadRequest, &resBody)
		return
	}

	resBody := models.NewSuccessResponse(http.StatusOK, map[string]interface{}{
		"token": token,
	})

	c.JSON(http.StatusOK, &resBody)
}

func RequireAuth(c *gin.Context) {
	token, err := c.Cookie("Authorization")
	if err != nil {
		resBody := models.NewFailedResponse(http.StatusUnauthorized, map[string]string{
			"error": "token is unvalid",
		})

		c.JSON(http.StatusUnauthorized, &resBody)
		return
	}

	if err := validateToken(token); err != nil {
		resBody := models.NewFailedResponse(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})

		c.JSON(http.StatusUnauthorized, &resBody)
		return
	}

	c.Next()
}

func generateJWTToken(c *gin.Context, username interface{}) (string, error) {
	claims := jwt.MapClaims{
		"name": username,
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(), // Token expiration time (30 days from now)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	return tokenString, nil
}

func validateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return fmt.Errorf("invalid token")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return fmt.Errorf("token expired")
	}

	var user models.User
	initializers.DB.Where(&models.User{
		Name: claims["name"].(string),
	}).First(&user)

	if user.ID == 0 {
		return fmt.Errorf("couldn't find user")
	}

	return nil
}
