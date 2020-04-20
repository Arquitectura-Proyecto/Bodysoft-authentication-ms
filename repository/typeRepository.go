package repository

import (
	"github.com/jpbmdev/Bodysoft-authentication-ms/data"
	"github.com/jpbmdev/Bodysoft-authentication-ms/models"
)

// GetTypes ..
func GetTypes(Type *[]models.Type) string {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Find(&Type).Error; err != nil {
		return err.Error()
	}
	return "ok"
}

// FindTypeByID ..
func FindTypeByID(Type models.Type) error {
	db := data.DatabaseConection()
	defer db.Close()
	if err := db.Where("ID = ?", Type.ID).First(&Type).Error; err != nil {
		return err
	}
	return nil
}
