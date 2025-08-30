package graphql

import (
	"libs/core/graphql/generated"
	"libs/core/graphql/resolver"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func NewGraphQLHandler(app *fiber.App) (err error) {
	srv := handler.New(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: &resolver.Resolver{}},
		),
	)
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.Use(extension.Introspection{})

	app.All("/graphql", adaptor.HTTPHandler(http.HandlerFunc(srv.ServeHTTP)))

	return
}
