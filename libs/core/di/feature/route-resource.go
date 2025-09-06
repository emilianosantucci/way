package feature

import (
	"libs/core/handler/rest"
	"libs/core/handler/rest/dto"
	"libs/core/handler/rest/handler"
	"libs/core/logging"
	"libs/core/model/convert"
	"libs/core/repository"
	"libs/core/service"

	"go.uber.org/fx"
)

var RouteResourceModule = fx.Module(featureResourcePrefix+"route",
	fx.Provide(
		fx.Private,
		repository.NewRouteResourceRepository,
		handler.NewRouteResourceHandler,
		convert.NewRouteResourceConverter,
		dto.NewRouteResourceConverter,
	),
	fx.Provide(
		service.NewRouteResourceService,
	),
	fx.Invoke(
		repository.RegisterRouterResourceEntities,
		rest.RegisterRouteResourceHandler,
	),
	fx.WithLogger(logging.FxLogger),
)
