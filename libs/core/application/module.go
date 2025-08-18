package application

import (
	"libs/core/application/handler"
	"libs/core/application/repository"
	"libs/core/application/service"

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
		handler.NewRest,
	),
	fx.Invoke(handler.RegisterRestHandler),
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log}
	}),
)
