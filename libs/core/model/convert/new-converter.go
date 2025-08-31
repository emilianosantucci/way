//go:build !goverter

package convert

import "libs/core/model/convert/generated"

func NewApplicationConverter() ApplicationConvert {
	return &generated.ApplicationConverter{}
}

func NewRestApiResourceConverter() RestApiResourceConvert {
	return &generated.RestApiResourceConverter{}
}

func NewRouteResourceConverter() RouteResourceConvert {
	return &generated.RouteResourceConverter{}
}
