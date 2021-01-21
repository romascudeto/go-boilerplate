package main

import (
	"fmt"
	Config "go-project/config"
	"go-project/middlewares"
	"go-project/routes"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	r := gin.New()
	r.Use(middlewares.BasicAuth())
	r = routes.SetupRouter(r)

	r.Run(":8000")
}
