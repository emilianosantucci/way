package graphql

func NewResolver(mutationResolver *MutationResolver, queryResolver *QueryResolver) (resolver *Resolver) {
	return &Resolver{
		query:    queryResolver,
		mutation: mutationResolver,
	}
}

type Resolver struct {
	query    *QueryResolver
	mutation *MutationResolver
}

func (r *Resolver) Query() *QueryResolver {
	return r.query
}

func (r *Resolver) Mutation() *MutationResolver {
	return r.mutation
}
