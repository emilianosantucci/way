package repository

import (
	"libs/core/entity"

	"gorm.io/gorm"
)

func RegisterApplicationEntities(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.Application{},
	)
}

func RegisterRestApiResourceEntities(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.RestApiResource{},
	)
}

func RegisterRouterResourceEntities(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.Route{},
	)
}
