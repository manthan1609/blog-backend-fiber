package database

import (
	"log"

	"github.com/manthan1609/todo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConnection *gorm.DB

func ConnectDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("database connection failed")
	}

	log.Println("database connection successfull")

	db.AutoMigrate(new(models.Blog))

	DBConnection = db
}
