package graphql

import (
	"context"
	"fmt"
	"libs/core/feature/resource/route/service"
	"libs/core/graphql/generated"
)

func NewQueryResolver(svc *service.Service) *QueryResolver {
	return &QueryResolver{
		service: svc,
	}
}

type QueryResolver struct {
	service *service.Service
}

// Routes is the resolver for the routes field.
func (r *QueryResolver) Routes(ctx context.Context) ([]*generated.Route, error) {
	panic(fmt.Errorf("not implemented: Routes - routes"))
}
