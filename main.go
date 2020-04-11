package main

import (
	"github.com/jpbmdev/Bodysoft-authentication-ms/controller"
	"github.com/jpbmdev/Bodysoft-authentication-ms/data"
)

func main() {
	data.DatabaseMigration()
	controller.HandleRequest()
}
