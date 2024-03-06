package controllers

import (
	"go.mod/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=root password=kwanpeemai0101 dbname=YogurtShop port=3306 sslmode=disable timezone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) // change the database provider if necessary

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Yogurt{}) // register Post model

	DB = database
}
