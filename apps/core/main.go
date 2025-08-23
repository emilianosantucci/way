package main

import (
	"context"
	"libs/core/common"
	"libs/core/database"
	"libs/core/environment"
	"libs/core/feature/application"
	"libs/core/feature/resource/rest"
	"libs/core/feature/resource/rest/repository"
	"libs/core/feature/resource/rest/repository/entity"
	"libs/core/logging"
	"libs/core/messaging"
	"libs/core/validation"
	"libs/core/web"

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

func TestingDI(log *zap.SugaredLogger, repo *repository.Repository) { // FIXME: remove me
	ctx := context.Background()
	var err error

	newRest := new(entity.Rest)

	newRest.Path = "/"
	newRest.Method = common.Get.String()

	if err = repo.Create(ctx, newRest); err != nil {
		log.Error(err)
		return
	}

	log.Debugf("Resource rest creation: %+v", newRest)

	updRest := new(entity.Rest)

	updRest.ID = newRest.ID
	updRest.Path = "/updated"
	updRest.Method = common.Post.String()

	if err = repo.Update(ctx, updRest); err != nil {
		log.Error(err)
		return
	}

	log.Debugf("Resource rest update: %+v", updRest)

	foundRest := new(entity.Rest)

	if foundRest, err = repo.FindById(ctx, updRest.ID); err != nil {
		log.Error(err)
		return
	}

	log.Debugf("Resource rest found: %+v", foundRest)

	if err = repo.Delete(ctx, updRest.ID); err != nil {
		log.Error(err)
		return
	}

	log.Debug("Resource rest deleted")

}
