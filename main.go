package main

import (
	"fmt"

	"./configs"
	"./models"
	"./routes"
	"github.com/jinzhu/gorm"
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