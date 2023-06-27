package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mufeedkka/goecommerce/database"
	"github.com/mufeedkka/goecommerce/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcom")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))
}
