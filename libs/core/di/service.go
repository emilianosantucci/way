package di

import (
	"libs/core/logging"
	"libs/core/service"

	"go.uber.org/fx"
)

var ServiceModule = fx.Module("service",
	fx.Provide(
		service.NewApplicationService,
		service.NewRestApiResourceService,
		service.NewRouteResourceService,
	),
	fx.WithLogger(logging.FxLogger),
)
