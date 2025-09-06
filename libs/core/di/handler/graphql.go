package handler

import (
	"libs/core/graphql"
	"libs/core/logging"

	"go.uber.org/fx"
)

var GraphQLHandlerModule = fx.Module("graphql-handler",
	fx.Invoke(graphql.RegisterGraphQLHandler),
	fx.WithLogger(logging.FxLogger),
)
