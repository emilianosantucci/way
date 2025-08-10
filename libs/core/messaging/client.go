package messaging

import (
	"context"
	"libs/core/configuration"

	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
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
}

func ClientLifecycle(params ClientLifecycleParams) {
	var (
		lc  = params.Lc
		nc  = params.NatsClient
		ns  = params.NatsServer
		err error
	)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if nc != nil && !nc.IsConnected() {
				nc, err = nats.Connect(ns.ClientURL())
			}
			return err
		},
		OnStop: func(ctx context.Context) error {
			if nc != nil && nc.IsConnected() {
				err = nc.Flush()
				nc.Close()
			}
			return err
		},
	})
}
