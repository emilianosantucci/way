package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Route struct {
	gorm.Model
	ID          uuid.UUID `gorm:"<-:create;primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Path        string    `gorm:"text;not null" json:"path"`
	Name        string    `gorm:"text;not null" json:"name"`
	Component   string    `gorm:"text;not null" json:"component"`
	Description string    `gorm:"text" json:"description"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
}

func (r *Route) TableName() string {
	return "resource_routes"
}
