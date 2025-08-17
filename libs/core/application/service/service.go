package service

import (
	"context"
	"errors"

	"libs/core/application/model"
	"libs/core/application/repository"
	"libs/core/common"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type Service struct {
	repository *repository.Repository
	validator  *validator.Validate
}

func NewService(repository *repository.Repository, validator *validator.Validate) *Service {
	return &Service{
		repository: repository,
		validator:  validator,
	}
}

func (s *Service) Create(ctx context.Context, newApp *model.NewApplication) (app *model.Application, err error) {
	if err = s.validator.StructCtx(ctx, newApp); err != nil {
		return
	}

	app = new(model.Application)
	if err = copier.Copy(&app, &newApp); err != nil {
		return
	}

	return app, s.repository.Create(ctx, app)
}

func (s *Service) Update(ctx context.Context, updApp *model.UpdateApplication) (app *model.Application, err error) {
	if err = s.validator.StructCtx(ctx, updApp); err != nil {
		return
	}

	app = new(model.Application)
	if err = copier.Copy(&app, &updApp); err != nil {
		return
	}

	if err = s.repository.Update(ctx, app); errors.Is(err, gorm.ErrDuplicatedKey) {
		err = common.ErrApplicationWithSameNameAndVersionExists
	}

	return
}

func (s *Service) FindById(ctx context.Context, id uuid.UUID) (app *model.Application, err error) {
	if err = s.validator.VarCtx(ctx, id, "uuid4_rfc4122"); err != nil {
		return
	}
	return s.repository.FindById(ctx, id)
}
