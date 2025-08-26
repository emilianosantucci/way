package model

import "libs/core/feature/resource/route/repository/entity"

// goverter:converter
// goverter:name Converter
// goverter:output:file ./converter.go
// goverter:ignoreMissing
// goverter:output:raw func NewConverter() Convert {
// goverter:output:raw    return &Converter{}
// goverter:output:raw }
type Convert interface {
	// goverter:update target
	FromNewToEntity(source *NewRoute, target *entity.Route)

	// goverter:update target
	FromUpdateToEntity(source *UpdateRoute, target *entity.Route)

	// goverter:update target
	ToModel(source *entity.Route, target *Route)
}
