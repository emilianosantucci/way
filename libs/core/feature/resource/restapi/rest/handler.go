package rest

import (
	"errors"
	"libs/core/common"
	"libs/core/feature/resource/restapi/mapper"
	"libs/core/feature/resource/restapi/model"
	"libs/core/feature/resource/restapi/rest/dto"
	"libs/core/feature/resource/restapi/service"

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
	app.Post("/resources/rest-apis/", handler.create)
	app.Put("/resources/rest-apis/:id", handler.update)
	app.Delete("/resources/rest-apis/:id", handler.delete)
	app.Get("/resources/rest-apis/:id", handler.findById)
}

type Handler struct {
	service   *service.Service
	log       *zap.SugaredLogger
	validator *validator.Validate
	mapper    mapper.RestDtoMap
}

func (r *Handler) create(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	request := new(dto.NewRestApiResource)

	if err = ctx.Bind().Body(request); err != nil {
		return
	}

	if err = r.validator.StructCtx(ctx, request); err != nil {
		return
	}

	newRestApi := new(model.NewRestApiResource)
	r.mapper.FromNewToModel(request, newRestApi)

	var app *model.RestApiResource
	if app, err = r.service.Create(ctx, newRestApi); err != nil {
		return
	}

	response := new(dto.RestApiResource)
	r.mapper.ToDto(app, response)

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (r *Handler) findById(ctx fiber.Ctx) (err error) {
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
	r.mapper.ToDto(restApi, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *Handler) update(ctx fiber.Ctx) (err error) {
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

	if err = r.mapper.FromUpdateToModel(request, updApp); err != nil {
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
	r.mapper.ToDto(app, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *Handler) delete(ctx fiber.Ctx) (err error) {
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
