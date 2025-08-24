package web

import (
	"libs/core/logging"

	"go.uber.org/fx"
)

var Module = fx.Module("web",
	fx.Provide(NewWeb),
	fx.Invoke(RunWeb),
	fx.WithLogger(logging.FxLogger),
)
