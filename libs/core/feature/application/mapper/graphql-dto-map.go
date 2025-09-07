package mapper

import (
	"libs/core/feature/application/model"
	"libs/core/graphql/generated"
)

// goverter:converter
// goverter:name GraphqlDtoMapper
// goverter:output:file ./graphql-dto-mapper.generated.go
// goverter:extend libs/core/common:UuidToString
// goverter:extend github.com/google/uuid:Parse
// goverter:ignoreMissing
type GraphqlDtoMap interface {
	// goverter:update target
	FromNewToModel(source *generated.NewApplication, target *model.NewApplication)

	// goverter:update target
	ToDto(source *model.Application, target *generated.Application)
}
