package main

import (
	"context"
	"fmt"
	"github.com/uptrace/bun/migrate"
	"go.uber.org/fx"
	"libs/core/configuration"
	"libs/migration"
)

func main() {
	app := fx.New(
		configuration.Module,
		migration.Module,
		fx.Invoke(RunMigrations),
	)
	app.Run()
}

func RunMigrations(migrator *migrate.Migrator) {
	var ctx = context.Background()

	migrator.Init(ctx)

	if err := migrator.Lock(ctx); err != nil {
		panic(err)
	}
	defer migrator.Unlock(ctx) //nolint:errcheck

	group, err := migrator.Migrate(ctx)
	if err != nil {
		panic(err)
	}
	if group.IsZero() {
		fmt.Printf("there are no new migrations to run (database is up to date)\n")
		return
	}
	fmt.Printf("migrated to %s\n", group)
	return
}
