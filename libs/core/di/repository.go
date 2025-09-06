package di

import (
	"libs/core/logging"
	"libs/core/repository"

	"go.uber.org/fx"
)

var RepositoryModule = fx.Module("repository",
	fx.Provide(
		repository.NewApplicationRepository,
		repository.NewRestApiResourceRepository,
		repository.NewRouteResourceRepository,
	),
	fx.Invoke(
		repository.RegisterEntities,
	),
	fx.WithLogger(logging.FxLogger),
)
