package restapi

import (
	"libs/core/feature/resource/restapi/mapper"
	"libs/core/feature/resource/restapi/repository"
	"libs/core/feature/resource/restapi/rest"
	"libs/core/feature/resource/restapi/service"
	"libs/core/logging"

	"go.uber.org/fx"
)

var Module = fx.Module("feature-resource-restapi",
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
