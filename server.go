package main

import (
	"log"

	"github.com/ashcoder666/g0-blog/database"
	"github.com/ashcoder666/g0-blog/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic("env load failed")
	}
	database.ConnectDB()
}
func main() {

	postgresDB, _ := database.DBconn.DB()

	defer postgresDB.Close()
	app := fiber.New()

	app.Use(logger.New())

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
