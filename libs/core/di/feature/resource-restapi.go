package feature

import (
	"libs/core/feature/resource/restapi/repository"
	"libs/core/feature/resource/restapi/service"
	"libs/core/feature/resource/restapi/service/model"
	"libs/core/handler/rest/handler"
	"libs/core/logging"

	"go.uber.org/fx"
)

var FeatureResourceRestApiModule = fx.Module("feature-resource-restapi",
	fx.Provide(
		fx.Private,
		model.NewConverter,
		repository.NewRepository,
		handler.NewRestApiResourceHandler,
	),
	fx.Provide(
		service.NewService,
	),
	fx.Invoke(
		repository.RegisterEntities,
	),
	fx.WithLogger(logging.FxLogger),
)
