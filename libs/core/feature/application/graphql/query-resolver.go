package graphql

import (
	"context"
	"libs/core/common"
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

func (r *QueryResolver) ApplicationsPaginated(ctx context.Context, size int, after *string, before *string) (pageRes *generated.ApplicationPagination, err error) {
	page := &common.CursorPageRequest{
		Size:   size,
		Before: before,
		After:  after,
	}

	var result *model.PaginatedApplications
	result, err = r.service.FindAllPaginated(ctx, page)

	if err != nil {
		return
	}

	pageRes = new(generated.ApplicationPagination)
	r.mapper.ToPagination(result, pageRes)

	return
}
