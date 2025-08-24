package database

import (
	"libs/core/logging"

	"go.uber.org/fx"
)

var Module = fx.Module("database",
	fx.Provide(NewDatabase),
	fx.WithLogger(logging.FxLogger),
)
