package route

import (
	"libs/core/feature/resource/route/mapper"
	"libs/core/feature/resource/route/repository"
	"libs/core/feature/resource/route/rest"
	"libs/core/feature/resource/route/service"
	"libs/core/logging"

	"go.uber.org/fx"
)

var Module = fx.Module("feature-resource-route",
	fx.Provide(
		fx.Private,
		mapper.NewModelMapper,
		repository.NewRepository,
		mapper.NewRestDtoMapper,
		rest.NewHandler,
	),
	fx.Provide(
		service.NewService,
	),
	fx.Invoke(
		repository.RegisterEntities,
		rest.RegisterHandler,
	),
	fx.WithLogger(logging.FxLogger),
)
