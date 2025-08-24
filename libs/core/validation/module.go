package validation

import (
	"libs/core/logging"

	"go.uber.org/fx"
)

var Module = fx.Module("configuration",
	fx.Provide(NewValidator),
	fx.Invoke(RegisterCustomHttpValidators),
	fx.WithLogger(logging.FxLogger),
)
