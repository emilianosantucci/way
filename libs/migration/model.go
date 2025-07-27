package migration

import (
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"libs/migration/migrations"
)

type MigratorParams struct {
	fx.In
	DB         *bun.DB
	Log        *zap.SugaredLogger
	Migrations *migrate.Migrations `optional:"true"`
}

type MigratorResult struct {
	fx.Out
	Migrator *migrate.Migrator
}

func NewMigrator(params MigratorParams) MigratorResult {
	return MigratorResult{Migrator: migrate.NewMigrator(params.DB, migrations.Migrations)}
}
