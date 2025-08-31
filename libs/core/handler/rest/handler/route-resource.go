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

func NewRouteResourceHandler(service *service.RouteResourceService, log *zap.SugaredLogger, validator *validator.Validate, converter dto2.RouteResourceConvert) *RouteResourceHandler {
	rest := &RouteResourceHandler{
		service,
		log,
		validator,
		converter,
	}

	return rest
}

type RouteResourceHandler struct {
	service   *service.RouteResourceService
	log       *zap.SugaredLogger
	validator *validator.Validate
	converter dto2.RouteResourceConvert
}

func (r *RouteResourceHandler) Create(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	request := new(dto2.NewRouteResource)

	if err = ctx.Bind().Body(request); err != nil {
		return
	}

	if err = r.validator.StructCtx(ctx, request); err != nil {
		return
	}

	newRoute := new(model2.NewRouteResource)
	r.converter.FromNewToModel(request, newRoute)

	var route *model2.RouteResource
	if route, err = r.service.Create(ctx, newRoute); err != nil {
		return
	}

	response := new(dto2.RouteResource)
	r.converter.ToDto(route, response)

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (r *RouteResourceHandler) FindById(ctx fiber.Ctx) (err error) {
	id := ctx.Params("id")
	if err = r.validator.VarCtx(ctx, id, "required,uuid4_rfc4122"); err != nil {
		return
	}

	var route *model2.RouteResource
	if route, err = r.service.FindById(ctx, uuid.MustParse(id)); err != nil {
		if errors.Is(err, common.ErrRouteNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	response := new(dto2.RouteResource)
	r.converter.ToDto(route, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *RouteResourceHandler) FindAll(ctx fiber.Ctx) (err error) {
	var routes []*model2.RouteResource
	if routes, err = r.service.FindAll(ctx); err != nil {
		return
	}

	response := make([]*dto2.RouteResource, len(routes))
	for i, route := range routes {
		response[i] = new(dto2.RouteResource)
		r.converter.ToDto(route, response[i])
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *RouteResourceHandler) FindByPath(ctx fiber.Ctx) (err error) {
	path := ctx.Params("path")
	if err = r.validator.VarCtx(ctx, path, "required,min=1"); err != nil {
		return
	}

	var route *model2.RouteResource
	if route, err = r.service.FindByPath(ctx, path); err != nil {
		if errors.Is(err, common.ErrRouteNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	response := new(dto2.RouteResource)
	r.converter.ToDto(route, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *RouteResourceHandler) FindByName(ctx fiber.Ctx) (err error) {
	name := ctx.Params("name")
	if err = r.validator.VarCtx(ctx, name, "required,min=1"); err != nil {
		return
	}

	var route *model2.RouteResource
	if route, err = r.service.FindByName(ctx, name); err != nil {
		if errors.Is(err, common.ErrRouteNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	response := new(dto2.RouteResource)
	r.converter.ToDto(route, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *RouteResourceHandler) Update(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	id := ctx.Params("id")
	request := new(dto2.UpdateRouteResource)
	if err = ctx.Bind().Body(request); err != nil {
		return
	}
	request.ID = id

	if err = r.validator.StructCtx(ctx, request); err != nil {
		return
	}

	updRoute := new(model2.UpdateRouteResource)

	if err = r.converter.FromUpdateToModel(request, updRoute); err != nil {
		return
	}

	var route *model2.RouteResource
	if route, err = r.service.Update(ctx, updRoute); err != nil {
		if errors.Is(err, common.ErrRouteNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	response := new(dto2.RouteResource)
	r.converter.ToDto(route, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *RouteResourceHandler) Delete(ctx fiber.Ctx) (err error) {
	id := ctx.Params("id")
	if err = r.validator.VarCtx(ctx, id, "required,uuid4_rfc4122"); err != nil {
		return
	}

	if err = r.service.Delete(ctx, uuid.MustParse(id)); err != nil {
		if errors.Is(err, common.ErrRouteNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
