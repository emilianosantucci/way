package configuration

import (
	"libs/core/database"
	"libs/core/logging"

	"go.uber.org/fx"
)

var DatabaseModule = fx.Module(configurationPrefix+"database",
	fx.Provide(database.NewDatabase),
	fx.WithLogger(logging.FxLogger),
)
