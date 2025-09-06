package feature

import (
	"libs/core/logging"
	"libs/core/model/convert"
	"libs/core/repository"
	"libs/core/rest"
	"libs/core/rest/dto"
	"libs/core/rest/handler"
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
