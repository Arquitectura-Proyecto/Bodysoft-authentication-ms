package data

import (
	"os"

	"github.com/jinzhu/gorm"
)

//DatabaseConection ..
//Development db "root:1234@tcp(host.docker.internal:3306)/authentication?charset=utf8&parseTime=True"
//Docker db "juanpablo:12345@tcp(boysoft-authentication-db:3306)/authentication?charset=utf8&parseTime=True"
func DatabaseConection() *gorm.DB {
	DBURL := os.Getenv("DBURL")
	DBPORT := os.Getenv("DBPORT")
	DBUSER := os.Getenv("DBUSER")
	DBPASSWORD := os.Getenv("DBPASSWORD")
	db, err := gorm.Open("mysql", DBUSER+":"+DBPASSWORD+"@tcp("+DBURL+":"+DBPORT+")/authentication?charset=utf8&parseTime=True")
	if err != nil {
		panic(err.Error())
	}
	return db
}
