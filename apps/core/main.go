package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"libs/core/configuration"
)

func main() {
	app := fx.New(
		configuration.Module,
		fx.Invoke(Run),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
	app.Run()
}

type Params struct {
	fx.In
	Log *zap.SugaredLogger
}

func Run(params Params) {
	params.Log.Info("Hello Way!!!")
}
