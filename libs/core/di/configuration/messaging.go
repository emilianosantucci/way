package configuration

import (
	"libs/core/logging"
	"libs/core/messaging"

	"go.uber.org/fx"
)

var MessagingModule = fx.Module("messaging",
	fx.Provide(
		messaging.NewServer,
		messaging.NewClient,
	),
	fx.Invoke(
		messaging.ServerLifecycle,
		messaging.ClientLifecycle,
	),
	fx.WithLogger(logging.FxLogger),
)
