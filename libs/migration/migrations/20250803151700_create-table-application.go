package migrations

import (
	"context"
	"github.com/uptrace/bun"
	"libs/core/application"
)

func init() {
	var (
		domain  = "application"
		logUp   = loggerUp.Named(domain)
		logDown = loggerDown.Named(domain)
	)

	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		logUp.Debug("Create table")
		_, err := db.NewCreateTable().
			Model((*application.Application)(nil)).
			IfNotExists().
			Exec(ctx)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		logDown.Debug("Drop table")
		_, err := db.NewDropTable().
			Model((*application.Application)(nil)).
			IfExists().
			Exec(ctx)
		return err
	})
}
