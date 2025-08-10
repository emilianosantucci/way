package configuration

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewWeb() *fiber.App {
	return fiber.New()
}

func RunWeb(log *zap.SugaredLogger, lc fx.Lifecycle, app *fiber.App, config Configuration) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) (err error) {
			go func() {
				log.Info("starting web server...")
				err = app.Listen(fmt.Sprintf("%s:%d", config.Web.Host, config.Web.Port))
			}()
			return err
		},
		OnStop: func(ctx context.Context) (err error) {
			log.Info("stopping web server...")
			return app.Shutdown()
		},
	})
}
