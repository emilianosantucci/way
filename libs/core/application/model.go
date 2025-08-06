package application

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	ID          string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name        string `gorm:"type:text;not null;uniqueIndex:name_and_version" json:"name"`
	Version     string `gorm:"type:text;not null;uniqueIndex:name_and_version" json:"version"`
	Description string `gorm:"type:text" json:"description"`
}
