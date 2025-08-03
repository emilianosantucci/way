package migrations

import (
	"context"
	"github.com/uptrace/bun"
)

func init() {
	var (
		domain  = "domain"
		logUp   = loggerUp.Named(domain)
		logDown = loggerDown.Named(domain)
	)
	
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		logUp.Debug("Up message")
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		logDown.Debug("Down message")
		return nil
	})
}
