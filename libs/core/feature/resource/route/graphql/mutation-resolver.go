package graphql

import (
	"context"
	"fmt"
	"libs/core/feature/resource/restapi/service"
	"libs/core/graphql/generated"
)

func NewMutationResolver(svc *service.Service) *MutationResolver {
	return &MutationResolver{
		service: svc,
	}
}

type MutationResolver struct {
	service *service.Service
}

// CreateRoute is the resolver for the createRoute field.
func (r *MutationResolver) CreateRoute(ctx context.Context, input generated.NewRoute) (*generated.Route, error) {
	panic(fmt.Errorf("not implemented: CreateRoute - createRoute"))
}
