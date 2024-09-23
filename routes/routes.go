package routes

import (
	"my-fiber-app/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// app.Post("/cashiers/:cashierId/login", controller.Login)
	// app.Get("/cashiers/:cashierId/logout", controller.Login)
	// app.Post("/cashiers/:cashierId/passcode", controller.Login)

	app.Get("/cashiers", controller.CahierList)
	app.Post("/cashiers", controller.CreateCashier)
	app.Get("/cashiers/:cashierId", controller.EditCashier)
	app.Get("/cashiers/:cashierId", controller.UpdateCashier)
	app.Get("/cashiers/:cashierId", controller.DeleteCashier)
}
