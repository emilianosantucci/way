package rest

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(app *fiber.App, handler *Handler) {
	app.Post("/applications", handler.Create)
	app.Get("/applications/:id", handler.FindById)
}
