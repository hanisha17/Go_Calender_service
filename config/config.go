package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"fmt"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:hell0R@azors1234@tcp(127.0.0.1:3306)/calender?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	

	fmt.Println("Database connection established")
}

func GetDB() *gorm.DB {
	return DB
}

func SetDB (db *gorm.DB) {
	 DB=db
}
