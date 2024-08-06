package database

import (
	"fmt"
	"log"
	"os"

	"github.com/manthan1609/todo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConnection *gorm.DB

func ConnectDB() {
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_name)

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
