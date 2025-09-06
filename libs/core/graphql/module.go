package graphql

import (
	"go.uber.org/fx"
)

var Module = fx.Module("graphql",
	fx.Provide(NewServer),
)
