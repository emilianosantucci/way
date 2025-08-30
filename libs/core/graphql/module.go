package graphql

import "go.uber.org/fx"

var Module = fx.Module("graphql",
	fx.Invoke(NewGraphQLHandler),
)
