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

func GetAllBlogs(c *fiber.Ctx) error {

	context := fiber.Map{
		"status": "ok",
		"msg":    "success",
	}

	var record []model.Blog

	database.DBconn.Find(&record)

	c.Status(200)
	context["data"] = record
	return c.JSON(context)
}

func GetSingleBlog(c *fiber.Ctx) error {
	context := fiber.Map{
		"status": "ok",
		"msg":    "success",
	}
	var record model.Blog

	id := c.Params("id")

	database.DBconn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not Found.")

		context["statusText"] = ""
		context["msg"] = "Record not Found."
		c.Status(400)
		return c.JSON(context)
	}

	context["data"] = record
	c.Status(200)
	return c.JSON(context)
}

func UpdateSingeleBlog(c *fiber.Ctx) error {

	context := fiber.Map{
		"status": "ok",
		"msg":    "success",
	}
	var record model.Blog

	id := c.Params("id")

	database.DBconn.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not Found.")

		context["statusText"] = ""
		context["msg"] = "Record not Found."
		c.Status(400)
		return c.JSON(context)
	}

	if err := c.BodyParser(&record); err != nil {
		log.Println("reading error")
	}

	database.DBconn.Save(record)

	context["data"] = record
	c.Status(200)
	return c.JSON(context)

}
