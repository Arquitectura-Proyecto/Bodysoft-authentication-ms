package repository

import (
	"github.com/jpbmdev/Bodysoft-authentication-ms/data"
	"github.com/jpbmdev/Bodysoft-authentication-ms/models"
)

// FindUserByEmail ..
func FindUserByEmail(user models.User) error {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("Email = ?", user.Email).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByEmailPass ..
func GetUserByEmailPass(user *models.User) error {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("Email = ? AND Password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByEmail ..
func GetUserByEmail(user *models.User) error {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("Email = ?", user.Email).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserUserByIDPass ..
func GetUserUserByIDPass(user *models.User) error {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("ID = ? AND Password = ?", user.ID, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser .
func CreateUser(user models.User) error {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser ..
func UpdateUser(user models.User) error {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// GenerateRecoveryEmail .
func GenerateRecoveryEmail(email string, password string) error {
	subject := "Recuperacion de Contraseña Bodysoft"
	htmlContent := "<h1> Su contraseña es: " + password + "</h1>"
	if err := data.SendEmail(email, htmlContent, subject); err != nil {
		return err
	}
	return nil
}
