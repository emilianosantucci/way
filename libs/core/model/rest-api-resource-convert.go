package model

import (
	"libs/core/entity"
)

// goverter:converter
// goverter:name RestApiResourceConverter
// goverter:output:file ./rest-api-resource-convert.generated.go
// goverter:ignoreMissing
// goverter:output:raw func NewRestApiResourceConverter() RestApiResourceConvert {
// goverter:output:raw    return &RestApiResourceConverter{}
// goverter:output:raw }
// goverter:extend libs/core/common/http:ToString
// goverter:extend libs/core/common/http:ToHttpMethod
type RestApiResourceConvert interface {
	// goverter:update target
	FromNewToEntity(source *NewRestApiResource, target *entity.RestApiResource)

	// goverter:update target
	FromUpdateToEntity(source *UpdateRestApiResource, target *entity.RestApiResource)

	// goverter:update target
	ToModel(source *entity.RestApiResource, target *RestApiResource)
}
