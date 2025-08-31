package web

import (
	"libs/core/logging"
	"libs/core/web/handler"

	"go.uber.org/fx"
)

var Module = fx.Module("web",
	fx.Provide(
		NewFiber,
	),
	fx.Invoke(
		RunFiber,
		handler.RegisterGraphQLHandler,
	),
	fx.WithLogger(logging.FxLogger),
)
