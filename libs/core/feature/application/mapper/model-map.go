package mapper

import (
	"libs/core/feature/application/entity"
	model2 "libs/core/feature/application/model"
)

// goverter:converter
// goverter:name ModelMapper
// goverter:output:file ./model-mapper.generated.go
// goverter:ignoreMissing
type ModelMap interface {
	// goverter:update target
	FromNewToEntity(source *model2.NewApplication, target *entity.Application)

	// goverter:update target
	ToModel(source *entity.Application, target *model2.Application)

	// goverter:update target
	FromUpdateToEntity(source *model2.UpdateApplication, target *entity.Application)
}
