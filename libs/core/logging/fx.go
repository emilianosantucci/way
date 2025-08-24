package logging

import (
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func FxLogger(log *zap.Logger) fxevent.Logger {
	return &fxevent.ZapLogger{Logger: log}
}
