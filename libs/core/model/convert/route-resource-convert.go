package convert

import (
	"libs/core/entity"
	"libs/core/model"
)

// goverter:converter
// goverter:name RouteResourceConverter
// goverter:output:file ./generated/route-resource-converter.go
// goverter:ignoreMissing
type RouteResourceConvert interface {
	// goverter:update target
	FromNewToEntity(source *model.NewRouteResource, target *entity.Route)

	// goverter:update target
	FromUpdateToEntity(source *model.UpdateRouteResource, target *entity.Route)

	// goverter:update target
	ToModel(source *entity.Route, target *model.RouteResource)
}
