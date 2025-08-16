package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	ID      uuid.UUID `gorm:"<-:create;primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name    string    `gorm:"<-:create;type:text;not null;uniqueIndex:name_and_version" json:"name"`
	Version string    `gorm:"type:text;not null;uniqueIndex:name_and_version" json:"version"`
}
