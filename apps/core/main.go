package main

import (
	"context"
	"libs/core/application"
	"libs/core/application/service"
	"libs/core/application/service/model"
	"libs/core/database"
	"libs/core/environment"
	"libs/core/logging"
	"libs/core/messaging"
	"libs/core/validation"
	"libs/core/web"

	"github.com/google/uuid"

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

func TestingDI(svc *service.Service, db *gorm.DB, ns *server.Server, nc *nats.Conn, lc fx.Lifecycle, log *zap.SugaredLogger) { // FIXME: remove me
	ctx := context.Background()

	var (
		sub *nats.Subscription
		err error
		app *model.Application
		id  uuid.UUID
	)

	app, err = svc.Create(ctx, &model.NewApplication{Name: "svc-app", Version: "1.0.0"})

	log.Debugf("App: %+v", app)

	if err != nil {
		log.Error(err.Error())
	}

	id, err = uuid.Parse("37a44bf5-8037-4a1a-9f83-f0ace1931554")

	if err != nil {
		log.Error(err.Error())
	}

	app, err = svc.FindById(ctx, id)

	log.Debugf("Find by id app: %+v", app)

	if err != nil {
		log.Error(err.Error())
	}

	app, err = svc.Update(ctx, &model.UpdateApplication{ID: id, Version: "1.1.0"})

	if err != nil {
		log.Error(err.Error())
	}

	log.Debugf("Updated app: %+v", app)

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
