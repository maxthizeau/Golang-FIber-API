package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/maxthizeau/api-fiber/database"
	"github.com/maxthizeau/api-fiber/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to my awesome API !")
}

func setupRoutes(app *fiber.App) {

	// welcome endpoint
	app.Get("/api", welcome)

	// User endpoint
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)

	// Product endpoints
	app.Get("/api/products/:id", routes.GetProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Post("/api/products", routes.CreateProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)

	// Order endpoints
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)
	app.Post("/api/orders", routes.CreateOrder)
}

func main() {

	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
