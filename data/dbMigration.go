package data

import (
	"fmt"

	"github.com/jpbmdev/Bodysoft-authentication-ms/models"
)

// DatabaseMigration ..
func DatabaseMigration() {
	db := DatabaseConection()
	db.AutoMigrate(&models.User{}, &models.Type{})
	Type := models.Type{Name: "Entrenador"}
	if err := db.Where("name = ?", "Entrenador").First(&Type).Error; err != nil {
		db.Create(&Type)
	}
	Type = models.Type{Name: "Usuario"}
	if err := db.Where("name = ?", "Usuario").First(&Type).Error; err != nil {
		db.Create(&Type)
	}
	fmt.Println("Data Succesfully Migrated")
	db.Close()
}
