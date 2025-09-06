package configuration

import (
	"libs/core/logging"

	"go.uber.org/fx"
)

var LoggingModule = fx.Module(configurationPrefix+"logging",
	fx.Provide(logging.NewLogger),
	fx.WithLogger(logging.FxLogger),
)
