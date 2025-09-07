package mapper

import (
	model2 "libs/core/feature/application/model"
	"libs/core/feature/application/rest/dto"
)

// goverter:converter
// goverter:name RestDtoMapper
// goverter:output:file ./rest-dto-mapper.generated.go
// goverter:extend libs/core/common:UuidToString
// goverter:extend github.com/google/uuid:Parse
// goverter:ignoreMissing
type RestDtoMap interface {
	// goverter:update target
	FromNewToModel(source *dto.NewApplication, target *model2.NewApplication)

	// goverter:update target
	ToDto(source *model2.Application, target *dto.Application)

	// goverter:update target
	FromUpdateToModel(source *dto.UpdateApplication, target *model2.UpdateApplication) (err error)
}
