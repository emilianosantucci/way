package mapper

import (
	"libs/core/feature/resource/route/entity"
	"libs/core/feature/resource/route/model"
)

// goverter:converter
// goverter:name ModelMapper
// goverter:output:file ./model-mapper.generated.go
// goverter:ignoreMissing
type ModelMap interface {
	// goverter:update target
	FromNewToEntity(source *model.NewRoute, target *entity.Route)

	// goverter:update target
	FromUpdateToEntity(source *model.UpdateRoute, target *entity.Route)

	// goverter:update target
	ToModel(source *entity.Route, target *model.Route)
}
