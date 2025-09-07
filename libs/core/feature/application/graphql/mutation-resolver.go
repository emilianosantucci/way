package graphql

import (
	"context"
	"libs/core/feature/application/mapper"
	"libs/core/feature/application/model"
	"libs/core/feature/application/service"
	"libs/core/graphql/generated"

	"github.com/go-playground/validator/v10"
)

func NewMutationResolver(svc *service.Service, validator *validator.Validate, mapper mapper.GraphqlDtoMap) (resolver *MutationResolver) {
	return &MutationResolver{
		service:   svc,
		validator: validator,
		mapper:    mapper,
	}
}

type MutationResolver struct {
	service   *service.Service
	validator *validator.Validate
	mapper    mapper.GraphqlDtoMap
}

func (r *MutationResolver) CreateApplication(ctx context.Context, input generated.NewApplication) (app *generated.Application, err error) {
	newApp := new(model.NewApplication)
	r.mapper.FromNewToModel(&input, newApp)
	result, err := r.service.Create(ctx, newApp)
	if err != nil {
		return
	}
	app = new(generated.Application)
	r.mapper.ToDto(result, app)
	return
}
