package restapi

import (
	"libs/core/feature/resource/restapi/api"
	"libs/core/feature/resource/restapi/api/dto"
	"libs/core/feature/resource/restapi/repository"
	"libs/core/feature/resource/restapi/service"
	"libs/core/feature/resource/restapi/service/model"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

var Module = fx.Module("feature-resource-restapi",
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
	fx.Provide(api.NewRest),
	fx.Provide(
		fx.Private,
		dto.NewConverter,
	),
	fx.Invoke(api.RegisterApiRest),
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger { return &fxevent.ZapLogger{Logger: log} }),
)
