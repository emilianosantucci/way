package configuration

import (
	"libs/core/logging"
	"libs/core/validation"

	"go.uber.org/fx"
)

var ValidationModule = fx.Module(configurationPrefix+"validation",
	fx.Provide(validation.NewValidator),
	fx.Invoke(validation.RegisterCustomHttpValidators),
	fx.WithLogger(logging.FxLogger),
)
