package application

import "go.uber.org/fx"

type Result struct {
	fx.Out
	App *Application
}

func New() Result {
	return Result{
		App: &Application{},
	}
}

type Application struct {
}

func (a *Application) Log(name string) string {
	result := "Application " + name
	return result
}
