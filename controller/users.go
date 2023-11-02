package controller

import (
	"log"

	"github.com/ashcoder666/g0-blog/database"
	"github.com/ashcoder666/g0-blog/model"
	"github.com/gofiber/fiber/v2"
)

func SignupUser(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "Add a Blog",
	}

	var statuss int

	record := new(model.Users)
	statuss = 201
	if err := c.BodyParser(record); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid Data"})
		// log.Fatal("invalid data")
		// statuss = 400
		// context["msg"] = "error in data format"
	}

	result := database.DBconn.Create(record)

	if result.Error != nil {
		log.Println("error in saving data")
	}

	// fmt.Println(record)
	context["msg"] = "added succesfully"

	c.Status(statuss)
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

func UpdateUser(c *fiber.Ctx) error {
	db := database.DBconn
	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "updated",
		// "data":       record,
	}

	id := c.Params("id")

	var record model.Users
	db.First(&record, id)
	if record.ID == 0 {
		log.Println("Record not Found.")

		context["statusText"] = ""
		context["msg"] = "Record not Found."
		c.Status(400)
		return c.JSON(context)
	}
	if err := c.BodyParser(&record); err != nil {
		log.Println("Error in parsing request.")

		context["msg"] = "Something went wrong."
		c.Status(400)
		return c.JSON(context)
	}

	// fmt.Println(record)

	db.Save(record)
	context["data"] = record
	c.Status(200)
	return c.JSON(context)

}

func DeleteAuser(c *fiber.Ctx) error {
	db := database.DBconn

	context := fiber.Map{
		"statusText": "Ok",
		"msg":        "updated",
		// "data":       record,
	}
	id := c.Params("id")

	var record model.Users

	db.First(&record, id)

	if record.ID == 0 {
		log.Println("Record not Found.")

		context["statusText"] = ""
		context["msg"] = "Record not Found."
		c.Status(400)
		return c.JSON(context)
	}

	result := db.Delete(record)

	if result.Error != nil {
		log.Println("error in saving data")
	}

	c.Status(200)
	return c.JSON(context)

}
