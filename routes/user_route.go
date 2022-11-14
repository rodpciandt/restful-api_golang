package routes

import (
	"fiber-mongo-api/controller"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	
	app.Post("/user", controller.CreateUser)
	app.Get("/user/:userId", controller.GetUser)
}