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

type MessagingServerResult struct {
	fx.Out
	NatsServer *server.Server
}

func NewMessagingServer() (MessagingServerResult, error) {
	opts := &server.Options{
		ServerName:      "embedded",
		JetStream:       true,
		JetStreamDomain: "way",
	}

	ns, err := server.NewServer(opts)

	if err != nil {
		return MessagingServerResult{}, err
	}

	ns.ConfigureLogger()

	return MessagingServerResult{
		NatsServer: ns,
	}, err
}

type RunMessagingParams struct {
	fx.In
	Lc         fx.Lifecycle
	NatsServer *server.Server
}

func RunMessaging(params RunMessagingParams) {
	var (
		lc = params.Lc
		ns = params.NatsServer
	)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go ns.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			ns.Shutdown()
			go ns.WaitForShutdown()
			return nil
		},
	})
}

type ConnectMessagingClientParams struct {
	fx.In
	NatsServer    *server.Server
	Configuration configuration.Configuration
}

type ConnectMessagingClientResult struct {
	fx.Out
	NatsClient *nats.Conn
}

func NewMessagingClient(params ConnectMessagingClientParams) (ConnectMessagingClientResult, error) {
	var (
		cfg = params.Configuration
		ns  = params.NatsServer
	)

	if ready := ns.ReadyForConnections(time.Duration(cfg.Messaging.ServerReadyTimeout) * time.Second); !ready {
		return ConnectMessagingClientResult{}, errors.New("nats server startup timeout reached")
	}

	nc, err := nats.Connect(ns.ClientURL())

	return ConnectMessagingClientResult{
		NatsClient: nc,
	}, err

}
