package graphql

import (
	"context"
	"fmt"
	"libs/core/graphql/generated"
)

type RouteResourceQueryResolver struct {
}

// Routes is the resolver for the routes field.
func (r *RouteResourceQueryResolver) Routes(ctx context.Context) ([]*generated.Route, error) {
	panic(fmt.Errorf("not implemented: Routes - routes"))
}
