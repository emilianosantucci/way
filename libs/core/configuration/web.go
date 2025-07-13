package configuration

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type WebParams struct {
	fx.In
	Configuration Configuration
}

type WebResult struct {
	fx.Out
	App *fiber.App
}

func NewWeb(params WebParams) WebResult {
	return WebResult{App: fiber.New()} //fiber.New()
}

type RunWebParams struct {
	fx.In
	Lc            fx.Lifecycle
	Log           *zap.SugaredLogger
	Configuration Configuration
	App           *fiber.App
}

func RunWeb(params RunWebParams) {
	log, lc, app, config := params.Log, params.Lc, params.App, params.Configuration

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			var err error = nil
			go func() {
				log.Info("starting web server...")
				err = app.Listen(fmt.Sprintf("%s:%d", config.Web.Host, config.Web.Port))

			}()
			return err
		},
		OnStop: func(ctx context.Context) error {
			log.Info("stopping web server...")
			return app.Shutdown()
		},
	})
}
