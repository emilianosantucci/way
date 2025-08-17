package application

import (
	"libs/core/application/model"
	"libs/core/application/repository"
	"libs/core/application/rest"
	"libs/core/application/service"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Module = fx.Module("application",
	fx.Invoke(MigrateDomain),
	fx.Provide(repository.NewRepository),
	fx.Provide(service.NewService),
	fx.Provide(rest.NewRest),
	fx.Invoke(rest.RegisterRoutes),
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log}
	}),
)

func MigrateDomain(db *gorm.DB) error {
	return db.AutoMigrate(&model.Application{})
}
