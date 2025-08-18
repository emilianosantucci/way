package handler

import (
	"errors"
	"libs/core/application/handler/dto"
	"libs/core/application/service"
	"libs/core/application/service/model"
	"libs/core/common"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type Rest struct {
	service   *service.Service
	log       *zap.SugaredLogger
	validator *validator.Validate
}

func NewRest(service *service.Service, log *zap.SugaredLogger, validator *validator.Validate) *Rest {
	rest := &Rest{
		service,
		log,
		validator,
	}

	return rest
}

func RegisterRestHandler(app *fiber.App, handler *Rest) {
	app.Post("/applications", handler.create)
	app.Get("/applications/:id", handler.findById)
}

func (h *Rest) create(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	request := new(dto.NewApplication)

	if err = ctx.Bind().Body(request); err != nil {
		return
	}

	if err = h.validator.StructCtx(ctx, request); err != nil {
		return
	}

	newApp := new(model.NewApplication)
	if err = copier.Copy(newApp, request); err != nil {
		return
	}

	var app *model.Application
	if app, err = h.service.Create(ctx, newApp); err != nil {
		return
	}

	return ctx.Status(fiber.StatusCreated).JSON(app)
}

func (h *Rest) findById(ctx fiber.Ctx) (err error) {
	id := ctx.Params("id")

	if err = h.validator.VarCtx(ctx, id, "required,uuid4_rfc4122"); err != nil {
		return
	}

	var app *model.Application
	if app, err = h.service.FindById(ctx, uuid.MustParse(id)); err != nil {
		if errors.Is(err, common.ErrApplicationNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(err)
		}
		return
	}

	return ctx.Status(fiber.StatusOK).JSON(app)
}
