package repository

import (
	"libs/core/entity"

	"gorm.io/gorm"
)

func RegisterEntities(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.Application{},
		&entity.RestApiResource{},
		&entity.Route{},
	)
}
