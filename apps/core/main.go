package main

import (
	"fmt"
	"go.uber.org/fx"
	"libs/core/application"
)

func main() {
	app := fx.New(
		application.Module,
		fx.Invoke(Run),
	)

	app.Run()
}

type Params struct {
	fx.In
	App *application.Application
}

func Run(params Params) {
	result := params.App.Log("Way")
	fmt.Println(result)
}
