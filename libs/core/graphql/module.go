package graphql

import (
	"libs/core/logging"

	"go.uber.org/fx"
)

var Module = fx.Module("graphql",
	fx.Provide(NewServer),
	fx.Invoke(RegisterGraphQLHandler),
	fx.WithLogger(logging.FxLogger),
)
