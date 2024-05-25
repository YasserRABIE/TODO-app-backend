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
		resBody := models.NewFailedResponse(http.StatusBadRequest, map[string]string{
			"error": "Invalid request! Please provide username, email, and password",
		})
		c.JSON(http.StatusBadRequest, &resBody)
		return
	}

	if err := database.CreateAccount(registerReq); err != nil {
		resBody := models.NewFailedResponse(http.StatusConflict, map[string]string{
			"error": err.Error(),
		})
		c.JSON(http.StatusConflict, &resBody)
		return
	}

	resBody := models.NewSuccessResponse(http.StatusCreated, map[string]interface{}{
		"message": "Account created successfully",
	})
	c.JSON(http.StatusCreated, &resBody)
}

func LoginHandler(c *gin.Context) {
	loginReq := &models.LoginRequest{}
	c.Header("Content-Type", "application/json")

	if err := c.BindJSON(loginReq); err != nil {
		resBody := models.NewFailedResponse(http.StatusBadRequest, map[string]string{
			"error": "Invalid request! Please provide username and password",
		})
		c.JSON(http.StatusBadRequest, &resBody)
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

	c.Set("userName", account.UserName)
	c.Next()
}
