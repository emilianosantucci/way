package configuration

import (
	"libs/core/graphql"
	"libs/core/logging"

	"go.uber.org/fx"
)

var GraphQLModule = fx.Module(configurationPrefix+"graphql",
	fx.Provide(graphql.NewServer),
	fx.WithLogger(logging.FxLogger),
)
