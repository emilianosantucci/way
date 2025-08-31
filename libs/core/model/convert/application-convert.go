package convert

import (
	"libs/core/entity"
	"libs/core/model"
)

// goverter:converter
// goverter:name ApplicationConverter
// goverter:output:file ./generated/application-converter.go
// goverter:ignoreMissing
type ApplicationConvert interface {
	// goverter:update target
	FromNewToEntity(source *model.NewApplication, target *entity.Application)

	// goverter:update target
	ToModel(source *entity.Application, target *model.Application)

	// goverter:update target
	FromUpdateToEntity(source *model.UpdateApplication, target *entity.Application)
}
