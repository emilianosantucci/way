package migration

import (
	"go.uber.org/fx"
)

var Module = fx.Module("migrations",
	fx.Provide(NewMigrator),
)
