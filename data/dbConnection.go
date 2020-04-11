package data

import (
	"github.com/jinzhu/gorm"
)

//DatabaseConection ..
//Development db "root:1234@tcp(localhost:3306)/authentication?charset=utf8&parseTime=True"
//Docker db "juanpablo:12345@tcp(boysoft-authentication-db:3306)/authentication?charset=utf8&parseTime=True"
func DatabaseConection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:1234@tcp(localhost:3306)/authentication?charset=utf8&parseTime=True")
	if err != nil {
		panic(err.Error())
	}
	return db
}
