package handler

import (
	"errors"
	"libs/core/common"
	dto2 "libs/core/handler/rest/dto"
	model2 "libs/core/model"
	"libs/core/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func NewRestApiResourceHandler(service *service.RestApiResourceService, log *zap.SugaredLogger, validator *validator.Validate, converter dto2.RestApiResourceConvert) *RestApiResourceHandler {
	rest := &RestApiResourceHandler{
		service,
		log,
		validator,
		converter,
	}

	return rest
}

type RestApiResourceHandler struct {
	service   *service.RestApiResourceService
	log       *zap.SugaredLogger
	validator *validator.Validate
	converter dto2.RestApiResourceConvert
}

func (r *RestApiResourceHandler) Create(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	request := new(dto2.NewRestApiResource)

	if err = ctx.Bind().Body(request); err != nil {
		return
	}

	if err = r.validator.StructCtx(ctx, request); err != nil {
		return
	}

	newRestApi := new(model2.NewRestApiResource)
	r.converter.FromNewToModel(request, newRestApi)

	var app *model2.RestApiResource
	if app, err = r.service.Create(ctx, newRestApi); err != nil {
		return
	}

	response := new(dto2.RestApiResource)
	r.converter.ToDto(app, response)

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (r *RestApiResourceHandler) FindById(ctx fiber.Ctx) (err error) {
	id := ctx.Params("id")
	if err = r.validator.VarCtx(ctx, id, "required,uuid4_rfc4122"); err != nil {
		return
	}

	var restApi *model2.RestApiResource
	if restApi, err = r.service.FindById(ctx, uuid.MustParse(id)); err != nil {
		if errors.Is(err, common.ErrRestApiResourceNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	response := new(dto2.RestApiResource)
	r.converter.ToDto(restApi, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *RestApiResourceHandler) Update(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	id := ctx.Params("id")
	request := new(dto2.UpdateRestApiResource)
	if err = ctx.Bind().Body(request); err != nil {
		return
	}
	request.ID = id

	if err = r.validator.StructCtx(ctx, request); err != nil {
		return
	}

	updApp := new(model2.UpdateRestApiResource)

	if err = r.converter.FromUpdateToModel(request, updApp); err != nil {
		return
	}

	var app *model2.RestApiResource
	if app, err = r.service.Update(ctx, updApp); err != nil {
		if errors.Is(err, common.ErrRestApiResourceNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	response := new(dto2.RestApiResource)
	r.converter.ToDto(app, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *RestApiResourceHandler) Delete(ctx fiber.Ctx) (err error) {
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
