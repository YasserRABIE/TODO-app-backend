package database

import (
	"errors"

	"github.com/YasserRABIE/authentication-porject/initializers"
	"github.com/YasserRABIE/authentication-porject/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateAccount(c *gin.Context, a *models.AccountRequest) error {
	exists, err := checkIfUserExists(initializers.DB, a.UserName, a.Email)
	if exists {
		return errors.New(err.Error())
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), 10)
	if err != nil {
		return errors.New("failed to hash the pass")
	}

	user := &models.User{
		UserName: a.UserName,
		Email:    a.Email,
		Password: string(hash),
	}

	initializers.DB.Create(&user)
	return nil
}

func checkIfUserExists(db *gorm.DB, username, email string) (bool, error) {
	var user models.User

	db.Where(&models.User{UserName: username}).First(&user)
	if user.UserName != "" {
		return true, errors.New("the user name is used before")
	}

	db.Where(&models.User{Email: email}).First(&user)
	return user.Email != "", errors.New("the email is used before")
}
