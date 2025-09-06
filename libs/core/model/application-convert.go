package model

import (
	"libs/core/entity"
)

// goverter:converter
// goverter:name ApplicationConverter
// goverter:output:file ./application-convert.generated.go
// goverter:ignoreMissing
// goverter:output:raw func NewApplicationConverter() ApplicationConvert {
// goverter:output:raw    return &ApplicationConverter{}
// goverter:output:raw }
type ApplicationConvert interface {
	// goverter:update target
	FromNewToEntity(source *NewApplication, target *entity.Application)

	// goverter:update target
	ToModel(source *entity.Application, target *Application)

	// goverter:update target
	FromUpdateToEntity(source *UpdateApplication, target *entity.Application)
}
