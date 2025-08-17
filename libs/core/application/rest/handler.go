package rest

import (
	"libs/core/application/model"
	"libs/core/application/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
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

func (r *Handler) Create(ctx fiber.Ctx) (err error) {
	ctx.Accepts(fiber.MIMEApplicationJSON)

	dto := new(NewApplication)
	err = ctx.Bind().Body(dto)
	if err != nil {
		return
	}

	err = r.validator.StructCtx(ctx, dto)
	if err != nil {
		return
	}

	newApp := new(model.NewApplication)
	err = copier.Copy(newApp, dto)
	if err != nil {
		return
	}

	app, err := r.service.Create(ctx, newApp)

	if err != nil {
		return
	}

	return ctx.Status(fiber.StatusCreated).JSON(app)
}
