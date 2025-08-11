package main

import (
	"context"
	"libs/core/application"
	"libs/core/database"
	"libs/core/environment"
	"libs/core/logging"
	"libs/core/messaging"
	"libs/core/validation"
	"libs/core/web"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	app := fx.New(
		environment.Module,
		logging.Module,
		validation.Module,
		database.Module,
		web.Module,
		messaging.Module,
		application.Module,
		fx.Invoke(TestingDI),
	)
	app.Run()
}

func TestingDI(db *gorm.DB, ns *server.Server, nc *nats.Conn, lc fx.Lifecycle, log *zap.Logger) { // FIXME: remove me
	ctx := context.Background()
	gorm.G[application.Application](db).Create(ctx, &application.Application{
		Name: "test",
	})

	var (
		sub *nats.Subscription
		err error
	)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("starting nats subscriber...")
			sub, err = nc.Subscribe("prova", func(msg *nats.Msg) {
				name := string(msg.Data)
				err = msg.Respond([]byte("hello, " + name))
			})
			return err
		},
		OnStop: func(ctx context.Context) error {
			log.Info("stopping nats subscriber...")
			if sub != nil {
				err = sub.Unsubscribe()
			}
			return err
		},
	})
}
