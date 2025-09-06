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
