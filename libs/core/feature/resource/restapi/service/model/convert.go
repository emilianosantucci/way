package model

import "libs/core/feature/resource/restapi/repository/entity"

// goverter:converter
// goverter:name Converter
// goverter:output:file ./converter.go
// goverter:ignoreMissing
// goverter:output:raw func NewConverter() Convert {
// goverter:output:raw    return &Converter{}
// goverter:output:raw }
// goverter:extend libs/core/common/http:ToString
// goverter:extend libs/core/common/http:ToHttpMethod
type Convert interface {
	// goverter:update target
	FromNewToEntity(source *NewRestApiResource, target *entity.RestApiResource)

	// goverter:update target
	FromUpdateToEntity(source *UpdateRestApiResource, target *entity.RestApiResource)

	// goverter:update target
	ToModel(source *entity.RestApiResource, target *RestApiResource)
}
