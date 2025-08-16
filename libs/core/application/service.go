package application

import (
	"context"
	"libs/core/application/model"

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

func (s *Service) Create(ctx context.Context, model *model.Application) (err error) {
	err = s.validator.StructCtx(ctx, model)
	if err != nil {
		return
	}
	return s.repository.Create(ctx, model)
}
