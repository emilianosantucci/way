package model

import "libs/core/feature/resource/rest/repository/entity"

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
	FromNewToEntity(source *NewRestResource, target *entity.RestResource)

	// goverter:update target
	FromUpdateToEntity(source *UpdateRestResource, target *entity.RestResource)

	// goverter:update target
	ToModel(source *entity.RestResource, target *RestResource)
}
