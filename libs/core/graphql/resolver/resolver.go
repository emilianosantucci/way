package resolver

import (
	application "libs/core/feature/application/graphql"
	route "libs/core/feature/resource/route/graphql"
	"libs/core/graphql/generated"
)

// It serves as dependency injection for your application, add any dependencies you require here.

func NewResolver(appResolver *application.Resolver, routeResolver *route.Resolver) (r *Resolver) {
	return &Resolver{
		&mutationResolver{
			applicationMR: appResolver.Mutation(),
			routeMR:       routeResolver.Mutation(),
		},
		&queryResolver{
			applicationQR: appResolver.Query(),
			routeQR:       routeResolver.Query(),
		},
	}
}

type Resolver struct {
	*mutationResolver
	*queryResolver
}

// Mutation returns exec.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return r.mutationResolver }

// Query returns exec.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return r.queryResolver }
