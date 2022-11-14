package main

import (
	"fiber-mongo-api/config"
	"fiber-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// runs db
	config.ConnectDB()


	// routes
	routes.UserRoute(app) 

	// ports
	app.Listen(":6000")
}