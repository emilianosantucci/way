package application

import (
	"libs/core/feature/application/api"
	"libs/core/feature/application/api/dto"
	"libs/core/feature/application/repository"
	"libs/core/feature/application/service"
	"libs/core/feature/application/service/model"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Module = fx.Module("application",
	fx.Provide(
		fx.Private,
		repository.NewRepository,
	),
	fx.Invoke(repository.RegisterEntity),
	fx.Provide(service.NewService),
	fx.Provide(
		fx.Private,
		model.NewConverter,
	),
	fx.Provide(
		fx.Private,
		api.NewRest,
	),
	fx.Provide(
		fx.Private,
		dto.NewConverter,
	),
	fx.Invoke(api.RegisterApiRest),
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log}
	}),
)
