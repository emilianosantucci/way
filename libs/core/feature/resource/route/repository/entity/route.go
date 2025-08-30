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
	Description string    `gorm:"text" json:"description"`
}

func (r *Route) TableName() string {
	return "resource_routes"
}
