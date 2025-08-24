package dto

import (
	"libs/core/feature/resource/restapi/service/model"
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
// goverter:extend libs/core/common/http:ToString
// goverter:extend libs/core/common/http:ToHttpMethod
type Convert interface {
	// goverter:update target
	FromNewToModel(source *NewRestApiResource, target *model.NewRestApiResource)

	// goverter:update target
	ToDto(source *model.RestApiResource, target *RestApiResource)

	// goverter:update target
	FromUpdateToModel(source *UpdateRestApiResource, target *model.UpdateRestApiResource) (err error)
}
