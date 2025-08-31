package convert

import (
	"libs/core/entity"
	"libs/core/model"
)

// goverter:converter
// goverter:name RestApiResourceConverter
// goverter:output:file ./generated/rest-api-resource-converter.go
// goverter:ignoreMissing
// goverter:extend libs/core/common/http:ToString
// goverter:extend libs/core/common/http:ToHttpMethod
type RestApiResourceConvert interface {
	// goverter:update target
	FromNewToEntity(source *model.NewRestApiResource, target *entity.RestApiResource)

	// goverter:update target
	FromUpdateToEntity(source *model.UpdateRestApiResource, target *entity.RestApiResource)

	// goverter:update target
	ToModel(source *entity.RestApiResource, target *model.RestApiResource)
}
