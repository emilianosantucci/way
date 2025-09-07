package mapper

import (
	model2 "libs/core/feature/resource/route/model"
	"libs/core/feature/resource/route/rest/dto"
)

// goverter:converter
// goverter:name RestDtoMapper
// goverter:output:file ./rest-dto-mapper.generated.go
// goverter:extend libs/core/common:UuidToString
// goverter:extend github.com/google/uuid:Parse
// goverter:ignoreMissing
type RestDtoMap interface {
	// goverter:update target
	FromNewToModel(source *dto.NewRoute, target *model2.NewRoute)

	// goverter:update target
	ToDto(source *model2.Route, target *dto.Route)

	// goverter:update target
	FromUpdateToModel(source *dto.UpdateRoute, target *model2.UpdateRoute) (err error)
}
