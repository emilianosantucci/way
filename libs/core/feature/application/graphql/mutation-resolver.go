package graphql

import (
	"context"
	"fmt"
	graphql1 "libs/core/graphql/generated"
)

type ApplicationMutationResolver struct {
}

func (r *ApplicationMutationResolver) CreateApplication(ctx context.Context, input graphql1.NewApplication) (*graphql1.Application, error) {
	panic(fmt.Errorf("not implemented: CreateApplication - createApplication"))
}
