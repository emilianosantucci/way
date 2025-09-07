package graphql

func NewResolver(query *QueryResolver, mutation *MutationResolver) *Resolver {
	return &Resolver{
		query:    query,
		mutation: mutation,
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
