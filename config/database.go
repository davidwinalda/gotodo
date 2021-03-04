package config

import (
	"fmt"
	"log"
	"todolist/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gotodo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&models.User{}, &models.Task{}, &models.Project{}, &models.Label{}, &models.Status{})
	fmt.Println("Connected to Database")

	DB = db
}
