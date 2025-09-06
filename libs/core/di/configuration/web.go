package configuration

import (
	"libs/core/graphql"
	"libs/core/logging"
	"libs/core/web"

	"go.uber.org/fx"
)

var WebModule = fx.Module(configurationPrefix+"web",
	fx.Provide(
		web.NewFiber,
	),
	fx.Invoke(
		web.RunFiber,
		graphql.RegisterGraphQLHandler,
	),
	fx.WithLogger(logging.FxLogger),
)
