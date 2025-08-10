package messaging

import (
	"context"
	"errors"
	"libs/core/configuration"
	"time"

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
