package database

import (
	"net/http"

	"github.com/YasserRABIE/authentication-porject/initializers"
	"github.com/YasserRABIE/authentication-porject/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(c *gin.Context, a *models.AccountRequest) {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), 10)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash the pass" + err.Error(),
		})
		return
	}

	user := &models.User{
		UserName: a.UserName,
		Password: string(hash),
	}

	initializers.DB.Create(&user)
}
