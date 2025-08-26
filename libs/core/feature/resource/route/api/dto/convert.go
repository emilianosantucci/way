package dto

import (
	"libs/core/feature/resource/route/service/model"
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
	FromNewToModel(source *NewRoute, target *model.NewRoute)

	// goverter:update target
	ToDto(source *model.Route, target *Route)

	// goverter:update target
	FromUpdateToModel(source *UpdateRoute, target *model.UpdateRoute) (err error)
}
