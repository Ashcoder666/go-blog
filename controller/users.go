package controller

import (
	"github.com/ashcoder666/g0-blog/database"
	"github.com/ashcoder666/g0-blog/model"
	"github.com/gofiber/fiber/v2"
)

func SignupUser(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a Blog",
	}

	c.Status(200)
	return c.JSON(context)
}

func GetAllUsers(c *fiber.Ctx) error {

	db := database.DBconn

	var allUsers []model.Users

	db.Find(&allUsers)

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a Blog",
		"data":       allUsers,
	}

	c.Status(200)
	return c.JSON(context)

}
