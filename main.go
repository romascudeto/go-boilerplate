package main

import (
	"fmt"
	Config "go-project/config"
	"go-project/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()

	r := routes.SetupRouter()

	r.Run(":8000")
}
