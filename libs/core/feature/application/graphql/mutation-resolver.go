package graphql

import (
	"context"
	"fmt"
	appServices "libs/core/feature/application/service"
	graphql1 "libs/core/graphql/generated"
)

func NewMutationResolver(svc *appServices.Service) (resolver *MutationResolver) {
	return &MutationResolver{
		Service: svc,
	}
}

type MutationResolver struct {
	Service *appServices.Service
}

func (r *MutationResolver) CreateApplication(ctx context.Context, input graphql1.NewApplication) (*graphql1.Application, error) {
	panic(fmt.Errorf("not implemented: CreateApplication - createApplication"))
}
