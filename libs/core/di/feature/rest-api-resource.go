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

var RestApiResourceModule = fx.Module(featureResourcePrefix+"rest-api",
	fx.Provide(
		fx.Private,
		repository.NewRestApiResourceRepository,
		convert.NewRestApiResourceConverter,
		handler.NewRestApiResourceHandler,
		dto.NewRestApiResourceConverter,
	),
	fx.Provide(
		service.NewRestApiResourceService,
	),
	fx.Invoke(
		repository.RegisterRestApiResourceEntities,
		rest.RegisterRestApiResourceHandler,
	),
	fx.WithLogger(logging.FxLogger),
)
