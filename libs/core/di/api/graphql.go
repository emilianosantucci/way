package api

import (
	"libs/core/graphql"

	"go.uber.org/fx"
)

var GraphQLModule = fx.Module("graphql",
	fx.Provide(graphql.NewServer),
)
