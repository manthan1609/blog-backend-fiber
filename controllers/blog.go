package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/manthan1609/todo/database"
	"github.com/manthan1609/todo/models"
)

// List All Blogs
func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog List",
	}

	var blogs []models.Blog

	database.DBConnection.Find(&blogs)

	context["blogs"] = blogs

	c.Status(200)
	return c.JSON(context)

}

// Add a Blog
func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "OK",
		"msg":        "Blog Added",
	}

	// var blog models.Blog
	blog := new(models.Blog)

	if err := c.BodyParser(&blog); err != nil {
		log.Println("error in parsing request")
		c.Status(400)
		context["statusText"] = "failed"
		context["msg"] = "error in parsing request"
		return c.JSON(context)
	}

	result := database.DBConnection.Create(blog)

	if result.Error != nil {
		log.Println("error in creating blog")
		c.Status(500)
		context["statusText"] = "failed"
		context["msg"] = "database error"
		return c.JSON(context)
	}

	c.Status(201)
	return c.JSON(context)
}

// Update a Blog
func BlogUpdate(c *fiber.Ctx) error {
	log.Println(c.Params("id"))

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog Updated",
	}

	c.Status(200)
	return c.JSON(context)
}

// Delete a Blog
func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog Deleted",
	}
	c.Status(200)
	return c.JSON(context)
}
