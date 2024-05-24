package database

import (
	"errors"

	"github.com/YasserRABIE/authentication-porject/initializers"
	"github.com/YasserRABIE/authentication-porject/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(c *gin.Context, a *models.AccountRequest) error {
	exists, err := checkIfUserExists(a.UserName, a.Email)
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

func GetAccount(l *models.LoginRequest) (*models.User, error) {
	var user = &models.User{}

	initializers.DB.Where("username= ?", l.UserName).First(&user)
	if user.ID == 0 {
		return nil, errors.New("please create an account before proceeding")
	}

	err := verifyPassword(user.Password, l.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func checkIfUserExists(username, email string) (bool, error) {
	var user models.User

	initializers.DB.Where(&models.User{UserName: username}).First(&user)
	if user.UserName != "" {
		return true, errors.New("the user name is used before")
	}

	initializers.DB.Where(&models.User{Email: email}).First(&user)
	return user.Email != "", errors.New("the email is used before")
}

func verifyPassword(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	); err != nil {
		return err
	}

	return nil
}
