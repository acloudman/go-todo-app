package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"go-todo-app/configs"
	"go-todo-app/models"
	"go-todo-app/routes"
)

var err error

func main() {

	configs.DB, err = gorm.Open("mysql", configs.DbURL(configs.BuildDBConfig()))

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer configs.DB.Close()
	configs.DB.AutoMigrate(&models.Todo{})

	r := Routes.SetupRouter()
	// running
	r.Run()
}