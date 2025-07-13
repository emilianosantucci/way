package configuration

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Module = fx.Module("configuration",
	fx.Provide(Read),
	fx.Provide(NewLogger),
	fx.Provide(NewDatabase),
	fx.Provide(NewWeb),
	fx.Invoke(RunWeb),
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log}
	}),
)
