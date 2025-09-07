package rest

import (
	"errors"
	"libs/core/common"
	"libs/core/feature/application/mapper"
	model2 "libs/core/feature/application/model"
	"libs/core/feature/application/rest/dto"
	"libs/core/feature/application/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func NewHandler(service *service.Service, log *zap.SugaredLogger, validator *validator.Validate, mapper mapper.RestDtoMap) (handler *Handler) {
	return &Handler{
		service,
		log,
		validator,
		mapper,
	}
}

func RegisterHandler(app *fiber.App, handler *Handler) {
	app.Post("/applications", handler.create)
	app.Put("/applications/:id", handler.update)
	app.Delete("/applications/:id", handler.delete)
	app.Get("/applications/:id", handler.findById)
}

type Handler struct {
	service   *service.Service
	log       *zap.SugaredLogger
	validator *validator.Validate
	mapper    mapper.RestDtoMap
}

func (r *Handler) create(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	request := new(dto.NewApplication)

	if err = ctx.Bind().Body(request); err != nil {
		return
	}

	if err = r.validator.StructCtx(ctx, request); err != nil {
		return
	}

	newApp := new(model2.NewApplication)
	r.mapper.FromNewToModel(request, newApp)

	var app *model2.Application
	if app, err = r.service.Create(ctx, newApp); err != nil {
		return
	}

	response := new(dto.Application)
	r.mapper.ToDto(app, response)

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (r *Handler) findById(ctx fiber.Ctx) (err error) {
	id := ctx.Params("id")
	if err = r.validator.VarCtx(ctx, id, "required,uuid4_rfc4122"); err != nil {
		return
	}

	var app *model2.Application
	if app, err = r.service.FindById(ctx, uuid.MustParse(id)); err != nil {
		if errors.Is(err, common.ErrApplicationNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	response := new(dto.Application)
	r.mapper.ToDto(app, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *Handler) update(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	id := ctx.Params("id")
	request := new(dto.UpdateApplication)
	if err = ctx.Bind().Body(request); err != nil {
		return
	}
	request.ID = id

	if err = r.validator.StructCtx(ctx, request); err != nil {
		return
	}

	updApp := new(model2.UpdateApplication)

	if err = r.mapper.FromUpdateToModel(request, updApp); err != nil {
		return
	}

	var app *model2.Application
	if app, err = r.service.Update(ctx, updApp); err != nil {
		if errors.Is(err, common.ErrApplicationNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	response := new(dto.Application)
	r.mapper.ToDto(app, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *Handler) delete(ctx fiber.Ctx) (err error) {
	id := ctx.Params("id")
	if err = r.validator.VarCtx(ctx, id, "required,uuid4_rfc4122"); err != nil {
		return
	}

	if err = r.service.Delete(ctx, uuid.MustParse(id)); err != nil {
		if errors.Is(err, common.ErrApplicationNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
