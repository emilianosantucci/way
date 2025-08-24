package main

import (
	"context"
	"libs/core/common/http"
	"libs/core/database"
	"libs/core/environment"
	"libs/core/feature/application"
	"libs/core/feature/resource/rest"
	"libs/core/feature/resource/rest/service"
	"libs/core/feature/resource/rest/service/model"
	"libs/core/logging"
	"libs/core/messaging"
	"libs/core/validation"
	"libs/core/web"

	"github.com/google/uuid"
	"go.uber.org/fx"
	"go.uber.org/zap"
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
		rest.Module,
		fx.Invoke(TestingDI),
	)
	app.Run()
}

func TestingDI(log *zap.SugaredLogger, svc *service.Service) { // FIXME: remove me
	ctx := context.Background()
	var err error

	newRest := new(model.NewRestResource)

	newRest.Path = "/"
	newRest.Method = http.Get

	var id uuid.UUID
	if result, err := svc.Create(ctx, newRest); err != nil {
		log.Error(err)
		return
	} else {
		id = result.ID
	}

	log.Debugf("Resource rest creation: %+v", newRest)

	updRest := new(model.UpdateRestResource)

	updRest.ID = id
	updRest.Path = "/updated"
	updRest.Method = http.Post

	if _, err := svc.Update(ctx, updRest); err != nil {
		log.Error(err)
		return
	}

	log.Debugf("Resource rest update: %+v", updRest)

	foundRest := new(model.RestResource)

	if foundRest, err = svc.FindById(ctx, id); err != nil {
		log.Error(err)
		return
	}

	log.Debugf("Resource rest found: %+v", foundRest)

	if err = svc.Delete(ctx, updRest.ID); err != nil {
		log.Error(err)
		return
	}

	log.Debug("Resource rest deleted")

}
