package mapper

import (
	model2 "libs/core/feature/resource/restapi/model"
	"libs/core/feature/resource/restapi/rest/dto"
)

// goverter:converter
// goverter:name RestDtoMapper
// goverter:output:file ./rest-dto-mapper.generated.go
// goverter:extend libs/core/common:UuidToString
// goverter:extend github.com/google/uuid:Parse
// goverter:ignoreMissing
// goverter:extend libs/core/common/http:ToString
// goverter:extend libs/core/common/http:ToHttpMethod
type RestDtoMap interface {
	// goverter:update target
	FromNewToModel(source *dto.NewRestApiResource, target *model2.NewRestApiResource)

	// goverter:update target
	ToDto(source *model2.RestApiResource, target *dto.RestApiResource)

	// goverter:update target
	FromUpdateToModel(source *dto.UpdateRestApiResource, target *model2.UpdateRestApiResource) (err error)
}
