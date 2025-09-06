package graphql

import (
	"libs/core/graphql/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func NewServer() (srv *handler.Server, err error) {
	srv = handler.New(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: &Resolver{}},
		),
	)

	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.Use(extension.Introspection{})

	return
}
