package repository

import (
	"strconv"

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

// GetUserByEmail ..
func GetUserByEmail(user *models.User) error {
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

// GetUserByID ..
func GetUserByID(user *models.User) error {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("ID = ?", user.ID).First(&user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByIDPass ..
func GetUserByIDPass(user *models.User) error {
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
	subject := "Recuperación de Contraseña Bodysoft"
	htmlContent := "<h1> Su contraseña es: " + password + "</h1>"
	if err := data.SendEmail(email, htmlContent, subject); err != nil {
		return err
	}
	return nil
}

// GenerateValidationEmail .
func GenerateValidationEmail(email string, vcode uint) error {
	subject := "Código de Verification BodySoft"
	htmlContent := "<h1> Su codigo es: " + strconv.Itoa(int(vcode)) + "</h1>"
	if err := data.SendEmail(email, htmlContent, subject); err != nil {
		return err
	}
	return nil
}
