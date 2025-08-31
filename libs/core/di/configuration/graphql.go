package configuration

import (
	"libs/core/graphql"
	"libs/core/logging"

	"go.uber.org/fx"
)

var GraphQLModule = fx.Module("graphql",
	fx.Provide(graphql.NewServer),
	fx.WithLogger(logging.FxLogger),
)
