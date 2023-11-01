package controller

import "github.com/gofiber/fiber/v2"

func SignupUser(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a Blog",
	}

	c.Status(200)
	return c.JSON(context)
}
