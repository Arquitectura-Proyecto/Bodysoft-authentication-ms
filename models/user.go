package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User ..
type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string
	TypeID   uint
	VCode    uint
	Check    bool
	Profile  bool
}
