package messaging

import (
	"context"

	"github.com/gofiber/fiber/v3/log"
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)
import "github.com/nats-io/nats-server/v2/server"

func NewClient(ns *server.Server) (nc *nats.Conn, err error) {
	log.Info("connect nats client...")
	nc, err = nats.Connect(ns.ClientURL())
	return
}

func ClientLifecycle(lc fx.Lifecycle, nc *nats.Conn, log *zap.SugaredLogger) {
	var err error
	lc.Append(
		fx.StopHook(func(ctx context.Context) error {
			log.Info("disconnect nats client...")

			if nc != nil && nc.IsConnected() {
				log.Debug("flushing nats client...")
				err = nc.Flush()

				log.Info("closing nats client...")
				nc.Close()
			}
			return err
		}),
	)
}
