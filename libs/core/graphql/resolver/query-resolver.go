package resolver

import (
	application "libs/core/feature/application/graphql"
	route "libs/core/feature/resource/route/graphql"
)

type applicationQR = application.QueryResolver
type routeQR = route.QueryResolver

type queryResolver struct {
	*applicationQR
	*routeQR
}
