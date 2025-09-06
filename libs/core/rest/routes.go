package rest

import (
	"libs/core/rest/handler"

	"github.com/gofiber/fiber/v3"
)

func RegisterApplicationHandler(app *fiber.App, handler *handler.ApplicationHandler) {
	app.Post("/applications", handler.Create)
	app.Put("/applications/:id", handler.Update)
	app.Delete("/applications/:id", handler.Delete)
	app.Get("/applications/:id", handler.FindById)
}

func RegisterRestApiResourceHandler(app *fiber.App, handler *handler.RestApiResourceHandler) {
	app.Post("/resources/rest-apis/", handler.Create)
	app.Put("/resources/rest-apis/:id", handler.Update)
	app.Delete("/resources/rest-apis/:id", handler.Delete)
	app.Get("/resources/rest-apis/:id", handler.FindById)
}

func RegisterRouteResourceHandler(app *fiber.App, handler *handler.RouteResourceHandler) {
	app.Post("/resources/routes/", handler.Create)
	app.Put("/resources/routes/:id", handler.Update)
	app.Delete("/resources/routes/:id", handler.Delete)
	app.Get("/resources/routes/:id", handler.FindById)
	app.Get("/resources/routes/", handler.FindAll)
	app.Get("/resources/routes/by-path/:path", handler.FindByPath)
	app.Get("/resources/routes/by-name/:name", handler.FindByName)
}
