package feature

import (
	"libs/core/feature/resource/route/repository"
	"libs/core/feature/resource/route/service"
	"libs/core/feature/resource/route/service/model"
	"libs/core/handler/rest/handler"
	"libs/core/logging"

	"go.uber.org/fx"
)

var FeatureResourceRestModule = fx.Module("feature-resource-route",
	fx.Provide(
		fx.Private,
		model.NewConverter,
		repository.NewRepository,
		handler.NewRouteResourceHandler,
	),
	fx.Provide(
		service.NewService,
	),
	fx.Invoke(
		repository.RegisterEntities,
	),
	fx.WithLogger(logging.FxLogger),
)
