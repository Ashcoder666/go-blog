package router

import (
	"github.com/ashcoder666/g0-blog/controller"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Post("/signup", controller.SignupUser)
	app.Get("/users", controller.GetAllUsers)
	app.Patch("/users/:id", controller.UpdateUser)
	app.Delete("/users/:id", controller.DeleteAuser)
	app.Post("/blog", controller.CreateBlog)
}
