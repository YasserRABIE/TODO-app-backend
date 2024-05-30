package database

import (
	"fmt"

	"github.com/YasserRABIE/authentication-porject/initializers"
	"github.com/YasserRABIE/authentication-porject/models"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(a *models.AccountRequest) error {
	exists, err := checkIfUserExists(a.Name, a.Email)
	if exists {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash the pass")
	}

	user := &models.User{
		Name:     a.Name,
		Email:    a.Email,
		Password: string(hash),
	}

	initializers.DB.Create(&user)
	return nil
}

func GetAccount(l *models.LoginRequest) (*models.User, error) {
	var user = &models.User{}

	initializers.DB.Where(&models.User{Email: l.Email}).First(&user)
	if user.ID == 0 {
		return nil, fmt.Errorf("incorrect name or password")
	}

	err := verifyPassword(user.Password, l.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func checkIfUserExists(name, email string) (bool, error) {
	var user models.User

	initializers.DB.Where("name= ?", name).First(&user)
	if user.Name != "" {
		return true, fmt.Errorf("the name is already taken")
	}

	initializers.DB.Where("email= ?", email).First(&user)
	if user.Email != "" {
		return true, fmt.Errorf("the email is already registered")
	}
	return false, nil
}

func verifyPassword(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	); err != nil {
		return fmt.Errorf("incorrect password")
	}

	return nil
}
