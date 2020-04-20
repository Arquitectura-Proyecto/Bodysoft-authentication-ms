package services

import (
	"github.com/jpbmdev/Bodysoft-authentication-ms/models"
	"github.com/jpbmdev/Bodysoft-authentication-ms/repository"
)

// GetTypesService ..
func GetTypesService() ([]models.Type, string) {
	Type := []models.Type{}
	if err := repository.GetTypes(&Type); err != "ok" {
		return Type, err
	}

	return Type, "ok"
}

// FindTypeByIDService ..
func FindTypeByIDService(typeid uint) string {
	var Type models.Type
	Type.ID = typeid
	if err := repository.FindTypeByID(Type); err != nil {
		return "El tipo de cuenta ingresado no existe"
	}
	return "ok"
}
