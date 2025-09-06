package web

import (
	"libs/core/logging"

	"go.uber.org/fx"
)

var Module = fx.Module("web",
	fx.Provide(
		NewFiber,
	),
	fx.Invoke(
		RunFiber,
	),
	fx.WithLogger(logging.FxLogger),
)
