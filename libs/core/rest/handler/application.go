package handler

import (
	"errors"
	"libs/core/common"
	model2 "libs/core/model"
	"libs/core/rest/dto"
	"libs/core/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func NewApplicationHandler(service *service.ApplicationService, log *zap.SugaredLogger, validator *validator.Validate, converter dto.ApplicationConvert) *ApplicationHandler {
	rest := &ApplicationHandler{
		service,
		log,
		validator,
		converter,
	}

	return rest
}

type ApplicationHandler struct {
	service   *service.ApplicationService
	log       *zap.SugaredLogger
	validator *validator.Validate
	converter dto.ApplicationConvert
}

func (r *ApplicationHandler) Create(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	request := new(dto.NewApplication)

	if err = ctx.Bind().Body(request); err != nil {
		return
	}

	if err = r.validator.StructCtx(ctx, request); err != nil {
		return
	}

	newApp := new(model2.NewApplication)
	r.converter.FromNewToModel(request, newApp)

	var app *model2.Application
	if app, err = r.service.Create(ctx, newApp); err != nil {
		return
	}

	response := new(dto.Application)
	r.converter.ToDto(app, response)

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (r *ApplicationHandler) FindById(ctx fiber.Ctx) (err error) {
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
	r.converter.ToDto(app, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *ApplicationHandler) Update(ctx fiber.Ctx) (err error) {
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

	if err = r.converter.FromUpdateToModel(request, updApp); err != nil {
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
	r.converter.ToDto(app, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *ApplicationHandler) Delete(ctx fiber.Ctx) (err error) {
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
