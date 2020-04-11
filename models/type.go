package models

import (
	"github.com/jinzhu/gorm"
)

//Type ..
type Type struct {
	gorm.Model
	Name  string `gorm:"unique;not null"`
	Users []User
}
