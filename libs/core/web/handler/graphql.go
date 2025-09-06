package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func RegisterGraphQLHandler(app *fiber.App, srv *handler.Server) {
	app.All("/graphql", adaptor.HTTPHandler(http.HandlerFunc(srv.ServeHTTP)))
}
