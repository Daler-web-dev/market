package main

import (
	"fmt"
	db "my-fiber-app/config"
	"my-fiber-app/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Printf("Go sales api course started...")
	db.Connect()

	app := fiber.New()

	// Middleware для выполнения перед каждым запросом
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("Middleware executed")
		return c.Next()
	})

	routes.Setup(app)

	app.Listen(":8080")

}
