package services

import (
	"appointment-booking-app/config"
	"appointment-booking-app/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func LoginAdmin(loginDetails *models.LoginDetails) (string, error) {
	var admin models.Admin
	if err := config.DB.Where("username = ?", loginDetails.Email).First(&admin).Error; err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(loginDetails.Password)); err != nil {
		return "", errors.New("incorrect password")
	}
	// Generate token (this is just a placeholder, implement JWT or another method)
	token := "some-generated-token"
	return token, nil
}
