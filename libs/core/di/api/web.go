package api

import (
	"libs/core/logging"
	"libs/core/web"
	"libs/core/web/handler"

	"go.uber.org/fx"
)

var WebModule = fx.Module("web",
	fx.Provide(
		web.NewFiber,
	),
	fx.Invoke(
		web.RunFiber,
		handler.RegisterGraphQLHandler,
	),
	fx.WithLogger(logging.FxLogger),
)
