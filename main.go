package main

import (
	"./config"
	"./models"
	"./routers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var err error

	config.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Post{})
	config.DB.AutoMigrate(&models.Classroom{})

	config.DB.Create(&models.User{
		Name:  "Benjamin",
		Role:  "admin",
		Email: "benjamin@gondange.com",
	})

	r := routers.SetupRouters()
	// running
	r.Run()
}
