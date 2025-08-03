package migrations

import (
	"github.com/uptrace/bun/migrate"
	"go.uber.org/zap"
)

var (
	logger, _  = zap.NewDevelopment()
	loggerUp   = logger.Named("migration").Named("up")
	loggerDown = logger.Named("migration").Named("down")
	Migrations = migrate.NewMigrations()
)

func init() {
	if err := Migrations.DiscoverCaller(); err != nil {
		panic(err)
	}
}
