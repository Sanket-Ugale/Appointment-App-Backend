package services

import (
	"appointment-booking-app/config"
	"appointment-booking-app/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(loginDetails *models.LoginDetails) (string, error) {
	var user models.User
	if err := config.DB.Where("email = ?", loginDetails.Email).First(&user).Error; err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password)); err != nil {
		return "", errors.New("incorrect password")
	}
	// Generate token (this is just a placeholder, implement JWT or another method)
	token := "some-generated-token"
	return token, nil
}

func GetUserById(id string) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(id string, updatedUser *models.User) error {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return err
	}
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(id string) error {
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
