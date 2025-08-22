package application

import (
	"libs/core/feature/application/api"
	"libs/core/feature/application/api/dto"
	"libs/core/feature/application/repository"
	"libs/core/feature/application/service"

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
		api.NewRest,
	),
	fx.Provide(
		fx.Private,
		fx.Annotate(dto.NewConverter, fx.As(new(dto.Convert))),
	),
	fx.Invoke(api.RegisterApiRest),
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log}
	}),
)
