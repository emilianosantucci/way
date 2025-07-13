package main

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"libs/core/configuration"
)

func main() {
	app := fx.New(
		configuration.Module,
		fx.Invoke(Run),
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
