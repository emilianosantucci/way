package mapper

import (
	"libs/core/feature/application/entity"
	"libs/core/feature/application/model"
)

// goverter:converter
// goverter:name ModelMapper
// goverter:output:file ./model-mapper.generated.go
// goverter:ignoreMissing
type ModelMap interface {
	// goverter:update target
	FromNewToEntity(source *model.NewApplication, target *entity.Application)

	// goverter:update target
	ToModel(source *entity.Application, target *model.Application)

	// goverter:update target
	FromUpdateToEntity(source *model.UpdateApplication, target *entity.Application)

	ToModels(source []entity.Application) (target []model.Application)
}
