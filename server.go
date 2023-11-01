package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	type User struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	app.Post("/signup", func(c *fiber.Ctx) error {
		var userData User

		if err := c.BodyParser(&userData); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid Data"})
		}

		fmt.Println(userData)

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "success", "data": userData})
	})

	app.Listen(":5002")
}
