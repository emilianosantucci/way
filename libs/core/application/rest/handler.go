package rest

import (
	"errors"
	"libs/core/application/model"
	"libs/core/application/service"
	"libs/core/common"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

type Handler struct {
	service   *service.Service
	log       *zap.SugaredLogger
	validator *validator.Validate
}

func NewRest(service *service.Service, log *zap.SugaredLogger, validator *validator.Validate) *Handler {
	return &Handler{
		service,
		log,
		validator,
	}
}

func (h *Handler) Create(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	dto := new(NewApplication)

	if err = ctx.Bind().Body(dto); err != nil {
		return
	}

	if err = h.validator.StructCtx(ctx, dto); err != nil {
		return
	}

	newApp := new(model.NewApplication)
	if err = copier.Copy(newApp, dto); err != nil {
		return
	}

	var app *model.Application
	if app, err = h.service.Create(ctx, newApp); err != nil {
		return
	}

	return ctx.Status(fiber.StatusCreated).JSON(app)
}

func (h *Handler) FindById(ctx fiber.Ctx) (err error) {
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
