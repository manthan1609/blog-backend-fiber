package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manthan1609/todo/controllers"
)

// Routing
func SetUpRoutes(app *fiber.App) {
	app.Get("/blogs", controllers.BlogList)
	app.Post("/blog", controllers.BlogCreate)
	app.Put("/blog/:id", controllers.BlogUpdate)
	app.Delete("/blog/:id", controllers.BlogDelete)
}
