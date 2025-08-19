package api

import (
	"errors"
	"libs/core/application/api/dto"
	"libs/core/application/service"
	"libs/core/application/service/model"
	"libs/core/common"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

func NewRest(service *service.Service, log *zap.SugaredLogger, validator *validator.Validate) *Rest {
	rest := &Rest{
		service,
		log,
		validator,
	}

	return rest
}

func RegisterApiRest(app *fiber.App, handler *Rest) {
	app.Post("/applications", handler.create)
	app.Put("/applications/:id", handler.update)
	app.Get("/applications/:id", handler.findById)
}

type Rest struct {
	service   *service.Service
	log       *zap.SugaredLogger
	validator *validator.Validate
}

func (r *Rest) create(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	request := new(dto.NewApplication)

	if err = ctx.Bind().Body(request); err != nil {
		return
	}

	if err = r.validator.StructCtx(ctx, request); err != nil {
		return
	}

	newApp := new(model.NewApplication)
	if err = copier.Copy(newApp, request); err != nil {
		return
	}

	var app *model.Application
	if app, err = r.service.Create(ctx, newApp); err != nil {
		return
	}

	return ctx.Status(fiber.StatusCreated).JSON(app)
}

func (r *Rest) findById(ctx fiber.Ctx) (err error) {
	id := ctx.Params("id")
	if err = r.validator.VarCtx(ctx, id, "required,uuid4_rfc4122"); err != nil {
		return
	}

	var app *model.Application
	if app, err = r.service.FindById(ctx, uuid.MustParse(id)); err != nil {
		if errors.Is(err, common.ErrApplicationNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	return ctx.Status(fiber.StatusOK).JSON(app)
}

func (r *Rest) update(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	id := ctx.Params("id")
	if err = r.validator.VarCtx(ctx, id, "required,uuid4_rfc4122"); err != nil {
		return
	}

	request := new(dto.UpdateApplication)
	if err = ctx.Bind().Body(request); err != nil {
		return
	}
	request.ID = id

	if err = r.validator.StructCtx(ctx, request); err != nil {
		return
	}

	updApp := new(model.UpdateApplication)
	if err = copier.Copy(updApp, request); err != nil {
		return
	}

	var app *model.Application
	if app, err = r.service.Update(ctx, updApp); err != nil {
		if errors.Is(err, common.ErrApplicationNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	return ctx.Status(fiber.StatusOK).JSON(app)
}
