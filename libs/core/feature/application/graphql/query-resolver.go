package graphql

import (
	"context"
	"fmt"
	"libs/core/feature/application/mapper"
	"libs/core/feature/application/model"
	"libs/core/feature/application/service"
	"libs/core/graphql/generated"
)

func NewQueryResolver(svc *service.Service, mapper mapper.GraphqlDtoMap) (resolver *QueryResolver) {
	return &QueryResolver{
		service: svc,
		mapper:  mapper,
	}
}

type QueryResolver struct {
	service *service.Service
	mapper  mapper.GraphqlDtoMap
}

// Applications is the resolver for the applications field.
func (r *QueryResolver) Applications(ctx context.Context) (apps []*generated.Application, err error) {
	var result []model.Application
	result, err = r.service.FindAll(ctx)

	if err != nil {
		return
	}
	apps = r.mapper.ToDtos(result)

	return
}

func (r *QueryResolver) ApplicationsPaginated(ctx context.Context, size int, after string, before string) (apps *generated.ApplicationPagination, err error) {
	panic(fmt.Errorf("not implemented: Applications - applicationsPaginated"))
}
