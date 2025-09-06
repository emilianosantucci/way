package graphql

import (
	"context"
	"fmt"
	"libs/core/graphql/generated"
)

type RouteResourceMutationResolver struct {
}

// CreateRoute is the resolver for the createRoute field.
func (r *RouteResourceMutationResolver) CreateRoute(ctx context.Context, input generated.NewRoute) (*generated.Route, error) {
	panic(fmt.Errorf("not implemented: CreateRoute - createRoute"))
}
