package dto

import (
	"libs/core/feature/application/service/model"
)

// goverter:converter
// goverter:name Converter
// goverter:output:file ./converter.go
// goverter:extend libs/core/common:UuidToString
// goverter:extend github.com/google/uuid:Parse
// goverter:ignoreMissing
// goverter:output:raw func NewConverter() Convert {
// goverter:output:raw    return &Converter{}
// goverter:output:raw }
type Convert interface {
	// goverter:update target
	ToModelNew(source *NewApplication, target *model.NewApplication)

	// goverter:update target
	ToDto(source *model.Application, target *Application)

	// goverter:update target
	ToModelUpdate(source *UpdateApplication, target *model.UpdateApplication) (err error)
}
