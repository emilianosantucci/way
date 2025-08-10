package messaging

import (
	"context"
	"errors"
	"libs/core/configuration"
	"time"

	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
)
import "github.com/nats-io/nats-server/v2/server"

type ServerParams struct {
	fx.In
	Configuration configuration.Configuration
}

type ServerResult struct {
	fx.Out
	NatsServer *server.Server
}

func NewServer(params ServerParams) (ServerResult, error) {
	var cfg = params.Configuration

	opts := &server.Options{
		ServerName:      "embedded",
		JetStream:       true,
		JetStreamDomain: "way",
	}

	ns, err := server.NewServer(opts)

	if err != nil {
		return ServerResult{}, err
	}

	ns.ConfigureLogger()
	ns.Start()

	if ready := ns.ReadyForConnections(time.Duration(cfg.Messaging.ServerReadyTimeout) * time.Second); !ready {
		return ServerResult{}, errors.New("nats server startup timeout reached")
	}

	return ServerResult{
		NatsServer: ns,
	}, err
}

type ServerLifecycleParams struct {
	fx.In
	Lc         fx.Lifecycle
	NatsServer *server.Server
}

func ServerLifecycle(params ServerLifecycleParams) {
	var (
		lc = params.Lc
		ns = params.NatsServer
	)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if ns != nil && !ns.Running() {
				ns.Start()
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if ns != nil && ns.Running() {
				ns.Shutdown()
				ns.WaitForShutdown()
			}
			return nil
		},
	})
}

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
