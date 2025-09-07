package mapper

import (
	"libs/core/feature/resource/restapi/entity"
	"libs/core/feature/resource/restapi/model"
)

// goverter:converter
// goverter:name ModelMapper
// goverter:output:file ./model-mapper.generated.go
// goverter:ignoreMissing
// goverter:extend libs/core/common/http:ToString
// goverter:extend libs/core/common/http:ToHttpMethod
type ModelMap interface {
	// goverter:update target
	FromNewToEntity(source *model.NewRestApiResource, target *entity.RestApiResource)

	// goverter:update target
	FromUpdateToEntity(source *model.UpdateRestApiResource, target *entity.RestApiResource)

	// goverter:update target
	ToModel(source *entity.RestApiResource, target *model.RestApiResource)
}
