package application

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	ID      uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name    string    `gorm:"type:text;not null;uniqueIndex:name_and_version" json:"name"`
	Version string    `gorm:"type:text;not null;uniqueIndex:name_and_version" json:"version"`
}

func (a Application) GetID() uuid.UUID {
	return a.ID
}
