package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"backend/models"
	"backend/utils"
	"fmt"
)

func ConnectDatabase() *gorm.DB {
	username := utils.Getenv("DB_USERNAME", "root")
	password := utils.Getenv("DB_PASSWORD", "Antimaling2@")
	database := utils.Getenv("DB_DATABASE", "db_ieee")
	host := utils.Getenv("DB_HOST", "127.0.0.1")

	dsn := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&models.Paper{})

	return db
}
