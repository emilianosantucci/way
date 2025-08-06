package main

import (
	"context"
	"libs/core/application"
	"libs/core/configuration"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

func main() {
	app := fx.New(
		configuration.Module,
		fx.Invoke(RunDB),
	)
	app.Run()
}

func RunDB(db *gorm.DB) { // FIXME: remove me
	err := db.AutoMigrate(&application.Application{})
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	gorm.G[application.Application](db).Create(ctx, &application.Application{
		Name: "test",
	})
}
