package messaging

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Module = fx.Module("messaging",
	fx.Provide(NewMessagingServer),
	fx.Provide(NewMessagingClient),
	fx.Invoke(RunMessaging),
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log}
	}),
)
