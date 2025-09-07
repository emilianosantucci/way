package resolver

import (
	application "libs/core/feature/application/graphql"
	route "libs/core/feature/resource/route/graphql"
)

type applicationMR = application.MutationResolver
type routeMR = route.MutationResolver

type mutationResolver struct {
	*applicationMR
	*routeMR
}
