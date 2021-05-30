package config

import (
	"fmt"
	"go-blog/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)
	dbCon, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("Failed to connect to database!")

	}

	dbCon.AutoMigrate(&models.Blog{})

	db = dbCon
}

func GetDB() *gorm.DB {
	return db
}
