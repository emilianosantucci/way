package messaging

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Module = fx.Module("messaging",
	fx.Provide(NewServer),
	fx.Invoke(ServerLifecycle),
	fx.Provide(NewClient),
	fx.Invoke(ClientLifecycle),
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log}
	}),
)
