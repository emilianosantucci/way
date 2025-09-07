package graphql

import (
	"context"
	"fmt"
	appServices "libs/core/feature/application/service"
	"libs/core/graphql/generated"
)

func NewQueryResolver(svc *appServices.Service) (resolver *QueryResolver) {
	return &QueryResolver{
		service: svc,
	}
}

type QueryResolver struct {
	service *appServices.Service
}

// Applications is the resolver for the applications field.
func (r *QueryResolver) Applications(ctx context.Context) ([]*generated.Application, error) {
	panic(fmt.Errorf("not implemented: Applications - applications"))
}
