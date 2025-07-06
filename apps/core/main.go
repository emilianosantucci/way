package main

import (
	"fmt"
	"go.uber.org/fx"
	"libs/core/application"
)

func main() {
	app := fx.New(
		application.Module,
		fx.Invoke(Hello),
	)

	app.Run()
}

func Hello(myApp *application.Application) {
	result := myApp.Log("Way")
	fmt.Println(result)
}
