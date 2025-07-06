package application

import "go.uber.org/fx"

var Module = fx.Module("application",
	fx.Provide(New),
)

func New() *Application {
	return &Application{}
}

type Application struct {
}

func (a *Application) Log(name string) string {
	result := "Application " + name
	return result
}
