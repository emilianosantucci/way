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

var ApplicationModule = fx.Module(featurePrefix+"application",
	fx.Provide(
		fx.Private,
		repository.NewApplicationRepository,
		convert.NewApplicationConverter,
		dto.NewApplicationConverter,
		handler.NewApplicationHandler,
	),
	fx.Provide(
		service.NewApplicationService,
	),
	fx.Invoke(
		repository.RegisterApplicationEntities,
		rest.RegisterApplicationHandler,
	),
	fx.WithLogger(logging.FxLogger),
)
