package messaging

import (
	"libs/core/logging"

	"go.uber.org/fx"
)

var Module = fx.Module("messaging",
	fx.Provide(
		NewServer,
		NewClient,
	),
	fx.Invoke(
		ServerLifecycle,
		ClientLifecycle,
	),
	fx.WithLogger(logging.FxLogger),
)
