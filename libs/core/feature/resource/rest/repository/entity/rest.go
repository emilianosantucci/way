package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RestResource struct {
	gorm.Model
	ID     uuid.UUID `gorm:"<-:create;primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Path   string    `gorm:"text;not null" json:"path"`
	Method string    `gorm:"text;not null" json:"method"`
}
