package route

import (
	"libs/core/feature/resource/route/api"
	"libs/core/feature/resource/route/api/dto"
	"libs/core/feature/resource/route/repository"
	"libs/core/feature/resource/route/service"
	"libs/core/feature/resource/route/service/model"
	"libs/core/logging"

	"go.uber.org/fx"
)

var Module = fx.Module("feature-resource-route",
	fx.Provide(
		fx.Private,
		model.NewConverter,
		repository.NewRepository,
		dto.NewConverter,
		api.NewRest,
	),
	fx.Provide(
		service.NewService,
	),
	fx.Invoke(
		repository.RegisterEntities,
		api.RegisterApiRest,
	),
	fx.WithLogger(logging.FxLogger),
)
