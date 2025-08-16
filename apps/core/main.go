package main

import (
	"context"
	"libs/core/application"
	"libs/core/application/model"
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

func TestingDI(repo *application.Repository, db *gorm.DB, ns *server.Server, nc *nats.Conn, lc fx.Lifecycle, log *zap.Logger) { // FIXME: remove me
	ctx := context.Background()

	var (
		sub *nats.Subscription
		err error
	)

	err = repo.Create(ctx, &model.Application{Name: "prova"})

	if err != nil {
		log.Error(err.Error())
	}

	err = repo.Update(ctx, "d65c285c-42ee-462c-a9ed-bd29842ce1f3", &model.Application{Name: "prova2"})

	if err != nil {
		log.Error(err.Error())
	}

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
