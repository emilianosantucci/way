package application

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Module("application",
	fx.Invoke(MigrateDomain),
)

func MigrateDomain(db *gorm.DB) error {
	return db.AutoMigrate(&Application{})
}
