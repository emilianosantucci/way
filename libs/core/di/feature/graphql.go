package feature

import (
	"libs/core/graphql"
	"libs/core/logging"

	"go.uber.org/fx"
)

var GraphQLSchemaModule = fx.Module(featurePrefix+"graphql-schema",
	fx.Invoke(graphql.RegisterGraphQLHandler),
	fx.WithLogger(logging.FxLogger),
)
