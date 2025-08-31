package feature

import (
	"libs/core/feature/resource/restapi/api"
	"libs/core/feature/resource/restapi/api/dto"
	"libs/core/feature/resource/restapi/repository"
	"libs/core/feature/resource/restapi/service"
	"libs/core/feature/resource/restapi/service/model"
	"libs/core/logging"

	"go.uber.org/fx"
)

var FeatureResourceRestApiModule = fx.Module("feature-resource-restapi",
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
