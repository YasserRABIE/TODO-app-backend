package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/YasserRABIE/authentication-porject/initializers"
	"github.com/YasserRABIE/authentication-porject/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func HandleAuth(c *gin.Context) {
	name, _ := c.Get("name")

	token, err := generateJWTToken(c, name)
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
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		c.Abort()
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
		c.Abort()
		return
	}

	token := parts[1]

	user, err := validateToken(token)
	if err != nil {
		resBody := models.NewFailedResponse(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})

		c.JSON(http.StatusUnauthorized, &resBody)
		c.Abort()
		return
	}

	c.Set("name", user.Name)
	c.Set("email", user.Email)
	c.Next()
}

func generateJWTToken(c *gin.Context, name interface{}) (string, error) {
	claims := jwt.MapClaims{
		"name": name,
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(), // Token expiration time (30 days from now)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie(
		"Authorization",
		tokenString,
		3600*24*30,
		"/",
		"https://todo-app-front-ixv5q2u4m-yasser-rabies-projects.vercel.app",
		true,
		true,
	)

	return tokenString, nil
}

func validateToken(tokenString string) (*models.User, error) {
	user := &models.User{}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, fmt.Errorf("token expired")
	}

	initializers.DB.Where(&models.User{
		Name: claims["name"].(string),
	}).First(&user)

	if user.ID == 0 {
		return nil, fmt.Errorf("couldn't find user")
	}

	return user, nil
}
