package handlers

import (
	"net/http"

	"github.com/YasserRABIE/authentication-porject/database"
	"github.com/YasserRABIE/authentication-porject/models"
	"github.com/gin-gonic/gin"
)

func HandleRegister(c *gin.Context) {
	registerReq := &models.AccountRequest{}
	c.Header("Content-Type", "application/json")

	if err := c.BindJSON(registerReq); err != nil {
		resBody := models.NewFailedResponse(400, map[string]string{
			"error": "Missing information! Please provide username and password",
		})

		c.JSON(http.StatusBadRequest, resBody)
		return
	}

	database.CreateAccount(c, registerReq)

	resBody := models.NewSuccessResponse(200, map[string]interface{}{
		"message": "account is added successfully",
	})

	c.JSON(http.StatusOK, resBody)
}
