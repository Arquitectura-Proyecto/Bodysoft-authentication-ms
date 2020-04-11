package data

import (
	"fmt"

	"github.com/jpbmdev/Bodysoft-authentication-ms/models"
)

// DatabaseMigration ..
func DatabaseMigration() {
	db := DatabaseConection()
	db.AutoMigrate(&models.User{}, &models.Type{})
	Type := models.Type{Name: "Trainer"}
	if err := db.Where("name = ?", "Trainer").First(&Type).Error; err != nil {
		db.Create(&Type)
	}
	Type = models.Type{Name: "Consumer"}
	if err := db.Where("name = ?", "Consumer").First(&Type).Error; err != nil {
		db.Create(&Type)
	}
	fmt.Println("Data Succesfully Migrated")
	db.Close()
}
