package api

import (
	"errors"
	"libs/core/common"
	"libs/core/feature/resource/restapi/api/dto"
	"libs/core/feature/resource/restapi/service"
	"libs/core/feature/resource/restapi/service/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func NewRest(service *service.Service, log *zap.SugaredLogger, validator *validator.Validate, converter dto.Convert) *Rest {
	rest := &Rest{
		service,
		log,
		validator,
		converter,
	}

	return rest
}

func RegisterApiRest(app *fiber.App, handler *Rest) {
	app.Post("/resources/rest-apis/", handler.create)
	app.Put("/resources/rest-apis/:id", handler.update)
	app.Delete("/resources/rest-apis/:id", handler.delete)
	app.Get("/resources/rest-apis/:id", handler.findById)
}

type Rest struct {
	service   *service.Service
	log       *zap.SugaredLogger
	validator *validator.Validate
	converter dto.Convert
}

func (r *Rest) create(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	request := new(dto.NewRestApiResource)

	if err = ctx.Bind().Body(request); err != nil {
		return
	}

	if err = r.validator.StructCtx(ctx, request); err != nil {
		return
	}

	newRestApi := new(model.NewRestApiResource)
	r.converter.FromNewToModel(request, newRestApi)

	var app *model.RestApiResource
	if app, err = r.service.Create(ctx, newRestApi); err != nil {
		return
	}

	response := new(dto.RestApiResource)
	r.converter.ToDto(app, response)

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (r *Rest) findById(ctx fiber.Ctx) (err error) {
	id := ctx.Params("id")
	if err = r.validator.VarCtx(ctx, id, "required,uuid4_rfc4122"); err != nil {
		return
	}

	var restApi *model.RestApiResource
	if restApi, err = r.service.FindById(ctx, uuid.MustParse(id)); err != nil {
		if errors.Is(err, common.ErrRestApiResourceNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	response := new(dto.RestApiResource)
	r.converter.ToDto(restApi, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *Rest) update(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	id := ctx.Params("id")
	request := new(dto.UpdateRestApiResource)
	if err = ctx.Bind().Body(request); err != nil {
		return
	}
	request.ID = id

	if err = r.validator.StructCtx(ctx, request); err != nil {
		return
	}

	updApp := new(model.UpdateRestApiResource)

	if err = r.converter.FromUpdateToModel(request, updApp); err != nil {
		return
	}

	var app *model.RestApiResource
	if app, err = r.service.Update(ctx, updApp); err != nil {
		if errors.Is(err, common.ErrRestApiResourceNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	response := new(dto.RestApiResource)
	r.converter.ToDto(app, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *Rest) delete(ctx fiber.Ctx) (err error) {
	id := ctx.Params("id")
	if err = r.validator.VarCtx(ctx, id, "required,uuid4_rfc4122"); err != nil {
		return
	}

	if err = r.service.Delete(ctx, uuid.MustParse(id)); err != nil {
		if errors.Is(err, common.ErrRestApiResourceNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
