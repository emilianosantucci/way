package application

import (
	"context"

	"github.com/emilianosantucci/way/core/application/model"
	"github.com/jinzhu/copier"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	repository *Repository
	validator  *validator.Validate
}

func NewService(repository *Repository, validator *validator.Validate) *Service {
	return &Service{
		repository: repository,
		validator:  validator,
	}
}

func (s *Service) Create(ctx context.Context, newApp *model.NewApplication) (app *model.Application, err error) {
	err = s.validator.StructCtx(ctx, newApp)
	if err != nil {
		return
	}
	app = new(model.Application)
	err = copier.Copy(&app, &newApp)
	if err != nil {
		return
	}
	return app, s.repository.Create(ctx, app)
}
