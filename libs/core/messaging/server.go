package messaging

import (
	"context"
	"errors"
	"time"

	"github.com/emilianosantucci/way/core/environment"

	"go.uber.org/fx"
)
import "github.com/nats-io/nats-server/v2/server"

func NewServer(config environment.Configuration) (ns *server.Server, err error) {
	opts := &server.Options{
		ServerName:      "embedded",
		JetStream:       true,
		JetStreamDomain: "way",
	}

	ns, err = server.NewServer(opts)

	if err != nil {
		return
	}

	ns.ConfigureLogger()
	ns.Start()

	if ready := ns.ReadyForConnections(time.Duration(config.Messaging.ServerReadyTimeout) * time.Second); !ready {
		return nil, errors.New("nats server startup timeout reached")
	}

	return
}

func ServerLifecycle(lc fx.Lifecycle, ns *server.Server) {
	lc.Append(
		fx.StopHook(func(ctx context.Context) (err error) {
			if ns != nil && ns.Running() {
				ns.Shutdown()
				ns.WaitForShutdown()
			}
			return
		}),
	)
}
