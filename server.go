package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"github.com/manthan1609/todo/database"
	"github.com/manthan1609/todo/router"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("error in loading .env file")
	}
	database.ConnectDB()
}

func main() {
	sqlDb, err := database.DBConnection.DB()

	if err != nil {
		panic("error in sql connection")
	}

	defer func() {
		sqlDb.Close()
		log.Println("database connection closed successfully")
	}()

	app := fiber.New()

	app.Use(logger.New())

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"message": "ok",
	// 	})

	// })

	router.SetUpRoutes(app)

	app.Listen(":8000")
}
