package main

import (
	"github.com/ashcoder666/g0-blog/database"
	"github.com/ashcoder666/g0-blog/router"
	"github.com/gofiber/fiber/v2"
)

func init() {
	database.ConnectDB()
}
func main() {

	postgresDB, _ := database.DBconn.DB()

	defer postgresDB.Close()
	app := fiber.New()

	// type User struct {
	// 	Name     string `json:"name"`
	// 	Email    string `json:"email"`
	// 	Password string `json:"password"`
	// }

	// app.Post("/signup", func(c *fiber.Ctx) error {
	// 	var userData User

	// 	if err := c.BodyParser(&userData); err != nil {
	// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid Data"})
	// 	}

	// 	fmt.Println(userData)

	// 	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "success", "data": userData})
	// })

	router.SetUpRoutes(app)

	app.Listen(":5002")
}
