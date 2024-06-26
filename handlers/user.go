package handlers

import (
	"net/http"

	"github.com/YasserRABIE/authentication-porject/database"
	"github.com/YasserRABIE/authentication-porject/models"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {
	registerReq := &models.AccountRequest{}
	c.Header("Content-Type", "application/json")

	if err := c.BindJSON(registerReq); err != nil {
		resBody := models.NewFailedResponse(http.StatusNoContent, map[string]string{
			"error": "Invalid request! Please provide name, email, and password",
		})

		c.JSON(http.StatusNoContent, &resBody)
		c.Abort()
		return
	}

	user, err := database.CreateAccount(registerReq)
	if err != nil {
		resBody := models.NewFailedResponse(http.StatusConflict, map[string]string{
			"error": err.Error(),
		})

		c.JSON(http.StatusConflict, &resBody)
		c.Abort()
		return
	}

	c.Set("name", user.Name)

	c.Next()
}

func LoginHandler(c *gin.Context) {
	loginReq := &models.LoginRequest{}
	c.Header("Content-Type", "application/json")

	if err := c.BindJSON(loginReq); err != nil {
		resBody := models.NewFailedResponse(http.StatusNoContent, map[string]string{
			"error": "Invalid request! Please provide name and password",
		})

		c.JSON(http.StatusNoContent, &resBody)
		c.Abort()
		return
	}

	account, err := database.GetAccount(loginReq)
	if err != nil {
		resBody := models.NewFailedResponse(http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
		c.JSON(http.StatusUnauthorized, &resBody)
		c.Abort()
		return
	}

	c.Set("name", account.Name)

	c.Next()
}

func GetAccountHandler(c *gin.Context) {
	name, _ := c.Get("name")
	email, _ := c.Get("email")

	resBody := models.NewSuccessResponse(http.StatusOK, map[string]interface{}{
		"name":  name,
		"email": email,
	})

	c.JSON(http.StatusOK, &resBody)

}
