package main

import (
	"go.uber.org/fx"
	"libs/core/configuration"
)

func main() {
	app := fx.New(
		configuration.Module,
	)
	app.Run()
}
