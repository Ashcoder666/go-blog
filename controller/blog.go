package controller

import (
	"log"

	"github.com/ashcoder666/g0-blog/database"
	"github.com/ashcoder666/g0-blog/model"
	"github.com/gofiber/fiber/v2"
)

func CreateBlog(c *fiber.Ctx) error {

	context := fiber.Map{
		"status": "ok",
		"msg":    "blog created succesfully",
	}

	record := new(model.Blog)

	if error := c.BodyParser(record); error != nil {
		return c.Status(400).JSON(fiber.Map{"msg": "invalid data type"})
	}

	result := database.DBconn.Create(record)

	if result.Error != nil {
		log.Fatal("something wrong in saving this to db")
	}

	context["data"] = record

	c.Status(200)

	return c.JSON(context)
}
