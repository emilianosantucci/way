package model

import "libs/core/feature/application/repository/entity"

// goverter:converter
// goverter:name Converter
// goverter:output:file ./converter.go
// goverter:ignoreMissing
// goverter:output:raw func NewConverter() Convert {
// goverter:output:raw    return &Converter{}
// goverter:output:raw }
type Convert interface {
	// goverter:update target
	FromNewToEntity(source *NewApplication, target *entity.Application)

	// goverter:update target
	ToModel(source *entity.Application, target *Application)

	// goverter:update target
	FromUpdateToEntity(source *UpdateApplication, target *entity.Application)
}
