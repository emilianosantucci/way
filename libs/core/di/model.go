package di

import (
	"libs/core/model/convert"

	"go.uber.org/fx"
)

var ModelModule = fx.Module("model",
	fx.Provide(
		convert.NewApplicationConverter,
		convert.NewRestApiResourceConverter,
		convert.NewRouteResourceConverter,
	),
)
