package dto

import (
	model2 "libs/core/model"
)

// goverter:converter
// goverter:name RouteResourceConverter
// goverter:output:file ./route-resource-convert.generated.go
// goverter:extend libs/core/common:UuidToString
// goverter:extend github.com/google/uuid:Parse
// goverter:ignoreMissing
// goverter:output:raw func NewRouteResourceConverter() RouteResourceConvert {
// goverter:output:raw    return &RouteResourceConverter{}
// goverter:output:raw }
type RouteResourceConvert interface {
	// goverter:update target
	FromNewToModel(source *NewRouteResource, target *model2.NewRouteResource)

	// goverter:update target
	ToDto(source *model2.RouteResource, target *RouteResource)

	// goverter:update target
	FromUpdateToModel(source *UpdateRouteResource, target *model2.UpdateRouteResource) (err error)
}
