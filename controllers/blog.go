package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// List All Blogs
func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog List",
	}

	c.Status(200)
	return c.JSON(context)

}

// Add a Blog
func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Blog Added",
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
