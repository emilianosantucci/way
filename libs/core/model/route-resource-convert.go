package model

import (
	"libs/core/entity"
)

// goverter:converter
// goverter:name RouteResourceConverter
// goverter:output:file ./route-resource-convert.generated.go
// goverter:ignoreMissing
// goverter:output:raw func NewRouteResourceConverter() RouteResourceConvert {
// goverter:output:raw    return &RouteResourceConverter{}
// goverter:output:raw }
type RouteResourceConvert interface {
	// goverter:update target
	FromNewToEntity(source *NewRoute, target *entity.Route)

	// goverter:update target
	FromUpdateToEntity(source *UpdateRoute, target *entity.Route)

	// goverter:update target
	ToModel(source *entity.Route, target *Route)
}
