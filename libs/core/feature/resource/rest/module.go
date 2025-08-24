package rest

import (
	"libs/core/feature/resource/rest/repository"
	"libs/core/feature/resource/rest/service"
	"libs/core/feature/resource/rest/service/model"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Module = fx.Module("feature-resource-rest",
	fx.Provide(
		fx.Private,
		repository.NewRepository,
	),
	fx.Invoke(repository.RegisterEntities),
	fx.Provide(service.NewService),
	fx.Provide(
		fx.Private,
		model.NewConverter,
	),
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger { return &fxevent.ZapLogger{Logger: log} }),
)
