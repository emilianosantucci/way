package application

import (
	"libs/core/feature/application/mapper"
	"libs/core/feature/application/repository"
	"libs/core/feature/application/rest"
	"libs/core/feature/application/service"
	"libs/core/logging"

	"go.uber.org/fx"
)

var Module = fx.Module("application",
	fx.Provide(
		fx.Private,
		repository.NewRepository,
		mapper.NewModelMapper,
		mapper.NewDtoMapper,
		rest.NewHandler,
	),
	fx.Invoke(
		repository.RegisterEntities,
		rest.RegisterHandler,
	),
	fx.Provide(
		service.NewService,
	),
	fx.WithLogger(logging.FxLogger),
)
