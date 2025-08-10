package messaging

import (
	"context"
	"libs/core/configuration"

	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
	"go.uber.org/zap"
)
import "github.com/nats-io/nats-server/v2/server"

type ClientParams struct {
	fx.In
	NatsServer    *server.Server
	Configuration configuration.Configuration
}

type ClientResult struct {
	fx.Out
	NatsClient *nats.Conn
}

func NewClient(params ClientParams) (ClientResult, error) {
	var ns = params.NatsServer

	nc, err := nats.Connect(ns.ClientURL())

	return ClientResult{
		NatsClient: nc,
	}, err
}

type ClientLifecycleParams struct {
	fx.In
	Lc         fx.Lifecycle
	NatsClient *nats.Conn
	NatsServer *server.Server
	Logger     *zap.SugaredLogger
}

func ClientLifecycle(params ClientLifecycleParams) {
	var (
		lc  = params.Lc
		nc  = params.NatsClient
		ns  = params.NatsServer
		log = params.Logger
		err error
	)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Start nats client...")

			if nc == nil || !nc.IsConnected() {
				log.Debug("Connecting nats client...")
				nc, err = nats.Connect(ns.ClientURL())
				if err == nil {
					log.Debug("Nats client connected!")
				}
			} else {
				log.Debug("Nats client exists and is connected!")
			}
			return err
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Disconnect nats client...")

			if nc != nil && nc.IsConnected() {
				log.Debug("Flushing nats client...")
				err = nc.Flush()

				log.Info("Closing nats client...")
				nc.Close()
			}
			return err
		},
	})
}
