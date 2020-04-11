package repository

import (
	"github.com/jpbmdev/Bodysoft-authentication-ms/data"
	"github.com/jpbmdev/Bodysoft-authentication-ms/models"
)

// FindUserByEmail ..
func FindUserByEmail(user models.User) string {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("Email = ?", user.Email).First(&user).Error; err != nil {
		return err.Error()
	}
	return "Usuario Encontrado"
}

// FindUserByEmailPass ..
func FindUserByEmailPass(user models.User) string {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("Email = ? AND Password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return err.Error()
	}
	return "Usuario Encontrado"
}

// FindUserByIDPass ..
func FindUserByIDPass(user models.User) string {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("ID = ? AND Password = ?", user.ID, user.Password).First(&user).Error; err != nil {
		return err.Error()
	}
	return "Usuario Encontrado"
}

// GetUserUserByEmail ..
func GetUserUserByEmail(user *models.User) string {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("Email = ?", user.Email).First(&user).Error; err != nil {
		return err.Error()
	}
	return "ok"
}

// GetUserUserByID ..
func GetUserUserByID(user *models.User) string {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("ID = ?", user.ID).First(&user).Error; err != nil {
		return err.Error()
	}
	return "ok"
}

// CreateUser .
func CreateUser(user models.User) string {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Create(&user).Error; err != nil {
		return err.Error()
	}
	return "ok"
}

// UpdateUser ..
func UpdateUser(user models.User) string {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Save(&user).Error; err != nil {
		return err.Error()
	}
	return "ok"
}

// GenerateEmail .
func GenerateEmail(email string, password string) string {
	if err := data.SendEmail(email, password); err != "ok" {
		return err
	}
	return "ok"
}
