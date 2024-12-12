package dao

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB() {
	dsn := "root:123456@tcp(8.148.22.219:3306)/library?charset=utf8mb4&parseTime=true&loc=local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connected")
}
