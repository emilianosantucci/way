package handler

import (
	"libs/core/handler/rest"
	"libs/core/handler/rest/dto"
	"libs/core/handler/rest/handler"

	"go.uber.org/fx"
)

var RestHandlerModule = fx.Module("rest-handler",
	fx.Provide(
		handler.NewApplicationHandler,
		handler.NewRestApiResourceHandler,
		handler.NewRouteResourceHandler,
		dto.NewApplicationConverter,
		dto.NewRestApiResourceConverter,
		dto.NewRouteResourceConverter,
	),
	fx.Invoke(
		rest.RegisterApplicationHandler,
		rest.RegisterRestApiResourceHandler,
		rest.RegisterRouteResourceHandler,
	),
)
