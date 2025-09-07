package graphql

import (
	"libs/core/graphql/resolver"
	"libs/core/logging"

	"go.uber.org/fx"
)

var Module = fx.Module("graphql",
	fx.Provide(
		resolver.NewResolver,
		NewServer,
	),
	fx.Invoke(RegisterGraphQLHandler),
	fx.WithLogger(logging.FxLogger),
)
