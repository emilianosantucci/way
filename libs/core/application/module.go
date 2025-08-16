package application

import (
	"github.com/emilianosantucci/way/core/application/model"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Module = fx.Module("application",
	fx.Invoke(MigrateDomain),
	fx.Provide(NewRepository),
	fx.Provide(NewService),
	fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		return &fxevent.ZapLogger{Logger: log}
	}),
)

func MigrateDomain(db *gorm.DB) error {
	return db.AutoMigrate(&model.Application{})
}
