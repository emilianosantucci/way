package configuration

import (
	"libs/core/environment"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var EnvironmentModule = fx.Module("environment",
	fx.Provide(environment.NewEnvironment),
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log}
	}),
)
