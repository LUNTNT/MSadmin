package main

import (
	"admin/DB"
	"admin/Router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	DB.MongoConnect()


	api := app.Group("/api")

	Router.CustomersRoute(api.Group("/admin"))

	app.Listen(":3002")

}