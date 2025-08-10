package main

import (
	"context"
	"libs/core/application"
	"libs/core/configuration"
	"libs/core/messaging"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

func main() {
	app := fx.New(
		configuration.Module,
		application.Module,
		messaging.Module,
		fx.Invoke(RunDB),
	)
	app.Run()
}

func RunDB(db *gorm.DB) { // FIXME: remove me
	ctx := context.Background()

	gorm.G[application.Application](db).Create(ctx, &application.Application{
		Name: "test",
	})
}
