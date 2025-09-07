package rest

import (
	"errors"
	"libs/core/common"
	"libs/core/feature/resource/route/mapper"
	model2 "libs/core/feature/resource/route/model"
	"libs/core/feature/resource/route/rest/dto"
	"libs/core/feature/resource/route/service"

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
	app.Post("/resources/routes/", handler.create)
	app.Put("/resources/routes/:id", handler.update)
	app.Delete("/resources/routes/:id", handler.delete)
	app.Get("/resources/routes/:id", handler.findById)
	app.Get("/resources/routes/", handler.findAll)
	app.Get("/resources/routes/by-path/:path", handler.findByPath)
	app.Get("/resources/routes/by-name/:name", handler.findByName)
}

type Handler struct {
	service   *service.Service
	log       *zap.SugaredLogger
	validator *validator.Validate
	mapper    mapper.RestDtoMap
}

func (r *Handler) create(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	request := new(dto.NewRoute)

	if err = ctx.Bind().Body(request); err != nil {
		return
	}

	if err = r.validator.StructCtx(ctx, request); err != nil {
		return
	}

	newRoute := new(model2.NewRoute)
	r.mapper.FromNewToModel(request, newRoute)

	var route *model2.Route
	if route, err = r.service.Create(ctx, newRoute); err != nil {
		return
	}

	response := new(dto.Route)
	r.mapper.ToDto(route, response)

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (r *Handler) findById(ctx fiber.Ctx) (err error) {
	id := ctx.Params("id")
	if err = r.validator.VarCtx(ctx, id, "required,uuid4_rfc4122"); err != nil {
		return
	}

	var route *model2.Route
	if route, err = r.service.FindById(ctx, uuid.MustParse(id)); err != nil {
		if errors.Is(err, common.ErrRouteNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	response := new(dto.Route)
	r.mapper.ToDto(route, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *Handler) findAll(ctx fiber.Ctx) (err error) {
	var routes []*model2.Route
	if routes, err = r.service.FindAll(ctx); err != nil {
		return
	}

	response := make([]*dto.Route, len(routes))
	for i, route := range routes {
		response[i] = new(dto.Route)
		r.mapper.ToDto(route, response[i])
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *Handler) findByPath(ctx fiber.Ctx) (err error) {
	path := ctx.Params("path")
	if err = r.validator.VarCtx(ctx, path, "required,min=1"); err != nil {
		return
	}

	var route *model2.Route
	if route, err = r.service.FindByPath(ctx, path); err != nil {
		if errors.Is(err, common.ErrRouteNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	response := new(dto.Route)
	r.mapper.ToDto(route, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *Handler) findByName(ctx fiber.Ctx) (err error) {
	name := ctx.Params("name")
	if err = r.validator.VarCtx(ctx, name, "required,min=1"); err != nil {
		return
	}

	var route *model2.Route
	if route, err = r.service.FindByName(ctx, name); err != nil {
		if errors.Is(err, common.ErrRouteNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	response := new(dto.Route)
	r.mapper.ToDto(route, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *Handler) update(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	id := ctx.Params("id")
	request := new(dto.UpdateRoute)
	if err = ctx.Bind().Body(request); err != nil {
		return
	}
	request.ID = id

	if err = r.validator.StructCtx(ctx, request); err != nil {
		return
	}

	updRoute := new(model2.UpdateRoute)

	if err = r.mapper.FromUpdateToModel(request, updRoute); err != nil {
		return
	}

	var route *model2.Route
	if route, err = r.service.Update(ctx, updRoute); err != nil {
		if errors.Is(err, common.ErrRouteNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	response := new(dto.Route)
	r.mapper.ToDto(route, response)

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (r *Handler) delete(ctx fiber.Ctx) (err error) {
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
