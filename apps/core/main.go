package main

import (
	"context"

	"github.com/emilianosantucci/way/core/application"
	"github.com/emilianosantucci/way/core/application/model"
	"github.com/emilianosantucci/way/core/database"
	"github.com/emilianosantucci/way/core/environment"
	"github.com/emilianosantucci/way/core/logging"
	"github.com/emilianosantucci/way/core/messaging"
	"github.com/emilianosantucci/way/core/validation"
	"github.com/emilianosantucci/way/core/web"

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

func TestingDI(svc *application.Service, db *gorm.DB, ns *server.Server, nc *nats.Conn, lc fx.Lifecycle, log *zap.SugaredLogger) { // FIXME: remove me
	ctx := context.Background()

	var (
		sub *nats.Subscription
		err error
		app *model.Application
	)

	app, err = svc.Create(ctx, &model.NewApplication{Name: "svc-app", Version: "1.0.0"})

	log.Debugf("App: %+v", app)

	if err != nil {
		log.Error(err.Error())
	}

	//err = repo.Update(ctx, "d65c285c-42ee-462c-a9ed-bd29842ce1f3", &model.Application{Name: "prova2"})
	//
	//if err != nil {
	//	log.Error(err.Error())
	//}

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
