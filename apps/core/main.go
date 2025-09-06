package main

import (
	"libs/core/di/api"
	"libs/core/di/configuration"
	"libs/core/di/feature"
	"libs/core/feature/resource/restapi/service"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	app := fx.New(
		api.WebModule,
		api.MessagingModule,
		api.GraphQLModule,
		configuration.EnvironmentModule,
		configuration.LoggingModule,
		configuration.DatabaseModule,
		configuration.ValidationModule,
		feature.ApplicationModule,
		feature.FeatureResourceRestApiModule,
		feature.FeatureResourceRestModule,
		fx.Invoke(TestingDI),
	)
	app.Run()
}

func TestingDI(log *zap.SugaredLogger, svc *service.Service, app *fiber.App) { // FIXME: remove me
	//ctx := context.Background()
	//var err error

	//newRest := new(model.NewRestApiResource)
	//
	//newRest.Path = "/"
	//newRest.Method = http.Get
	//
	//var id uuid.UUID
	//if result, err := svc.Create(ctx, newRest); err != nil {
	//	log.Error(err)
	//	return
	//} else {
	//	id = result.ID
	//}
	//
	//log.Debugf("Resource rest creation: %+v", newRest)
	//
	//updRest := new(model.UpdateRestApiResource)
	//
	//updRest.ID = id
	//updRest.Path = "/updated"
	//updRest.Method = http.All
	//
	//if _, err := svc.Update(ctx, updRest); err != nil {
	//	log.Error(err)
	//	return
	//}
	//
	//log.Debugf("Resource rest update: %+v", updRest)
	//
	//foundRest := new(model.RestApiResource)
	//
	//if foundRest, err = svc.FindById(ctx, id); err != nil {
	//	log.Error(err)
	//	return
	//}
	//
	//log.Debugf("Resource rest found: %+v", foundRest)
	//
	//if err = svc.Delete(ctx, updRest.ID); err != nil {
	//	log.Error(err)
	//	return
	//}
	//
	//log.Debug("Resource rest deleted")

}
