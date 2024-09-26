package routes

import (
	"my-fiber-app/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/cashiers/:cashierId/login", controller.Login)
	// app.Get("/cashiers/:cashierId/logout", controller.Login)
	// app.Get("/cashiers/:cashierId/passcode", controller.Login)

	app.Get("/cashiers", controller.CahierList)
	app.Get("/cashiers/:cashierId", controller.GetCashierDetails)
	app.Post("/cashiers", controller.CreateCashier)
	app.Patch("/cashiers/:cashierId", controller.UpdateCashier)
	app.Delete("/cashiers/:cashierId", controller.DeleteCashier)
}
