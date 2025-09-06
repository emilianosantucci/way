package dto

import (
	"libs/core/feature/resource/restapi/service/model"
)

// goverter:converter
// goverter:name RestApiResourceConverter
// goverter:output:file ./rest-api-resource-convert.generated.go
// goverter:extend libs/core/common:UuidToString
// goverter:extend github.com/google/uuid:Parse
// goverter:ignoreMissing
// goverter:output:raw func NewRestApiResourceConverter() RestApiResourceConvert {
// goverter:output:raw    return &RestApiResourceConverter{}
// goverter:output:raw }
// goverter:extend libs/core/common/http:ToString
// goverter:extend libs/core/common/http:ToHttpMethod
type RestApiResourceConvert interface {
	// goverter:update target
	FromNewToModel(source *NewRestApiResource, target *model.NewRestApiResource)

	// goverter:update target
	ToDto(source *model.RestApiResource, target *RestApiResource)

	// goverter:update target
	FromUpdateToModel(source *UpdateRestApiResource, target *model.UpdateRestApiResource) (err error)
}
