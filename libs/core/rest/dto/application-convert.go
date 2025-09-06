package dto

import (
	model2 "libs/core/model"
)

// goverter:converter
// goverter:name ApplicationConverter
// goverter:output:file ./application-convert.generated.go
// goverter:extend libs/core/common:UuidToString
// goverter:extend github.com/google/uuid:Parse
// goverter:ignoreMissing
// goverter:output:raw func NewApplicationConverter() ApplicationConvert {
// goverter:output:raw    return &ApplicationConverter{}
// goverter:output:raw }
type ApplicationConvert interface {
	// goverter:update target
	FromNewToModel(source *NewApplication, target *model2.NewApplication)

	// goverter:update target
	ToDto(source *model2.Application, target *Application)

	// goverter:update target
	FromUpdateToModel(source *UpdateApplication, target *model2.UpdateApplication) (err error)
}
