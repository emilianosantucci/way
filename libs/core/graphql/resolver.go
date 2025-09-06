package graphql

import (
	application "libs/core/feature/application/graphql"
	route "libs/core/feature/resource/route/graphql"
	"libs/core/graphql/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your application, add any dependencies you require here.

type Resolver struct{}

// Mutation returns exec.ApplicationMutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{} }

// Query returns exec.ApplicationQueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{} }

type mutationResolver struct {
	*application.ApplicationMutationResolver
	*route.RouteResourceMutationResolver
}
type queryResolver struct {
	*application.ApplicationQueryResolver
	*route.RouteResourceQueryResolver
}
