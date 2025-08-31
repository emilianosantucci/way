package feature

import (
	"libs/core/feature/application/repository"
	"libs/core/feature/application/service"
	"libs/core/feature/application/service/model"
	"libs/core/logging"

	"go.uber.org/fx"
)

var ApplicationModule = fx.Module("feature-application",
	fx.Provide(
		fx.Private,
		repository.NewRepository,
		model.NewConverter,
	),
	fx.Invoke(
		repository.RegisterEntities,
	),
	fx.Provide(
		service.NewService,
	),
	fx.WithLogger(logging.FxLogger),
)
