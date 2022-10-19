package database

import (
	"log"

	"github.com/GinAPIRest-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	dsn := "root:root@tcp(127.0.0.1:33060)/Students?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("error connecting to database")
	}
	DB.AutoMigrate(&models.Student{})
}
