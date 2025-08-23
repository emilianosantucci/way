package service

import (
	"context"
	"errors"
	"libs/core/common"
	"libs/core/feature/application/repository"
	"libs/core/feature/application/repository/entity"
	"libs/core/feature/application/service/model"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewService(repository *repository.Repository, validator *validator.Validate, converter model.Convert) *Service {
	return &Service{
		repository: repository,
		validator:  validator,
		converter:  converter,
	}
}

type Service struct {
	repository *repository.Repository
	validator  *validator.Validate
	converter  model.Convert
}

func (s *Service) Create(ctx context.Context, newApp *model.NewApplication) (app *model.Application, err error) {
	if err = s.validator.StructCtx(ctx, newApp); err != nil {
		return
	}

	ent := new(entity.Application)
	s.converter.FromNewToEntity(newApp, ent)

	if err = s.repository.Create(ctx, ent); err != nil {
		return
	}

	app = new(model.Application)
	s.converter.ToModel(ent, app)

	return
}

func (s *Service) Update(ctx context.Context, updApp *model.UpdateApplication) (app *model.Application, err error) {
	if err = s.validator.StructCtx(ctx, updApp); err != nil {
		return
	}

	ent := new(entity.Application)
	s.converter.FromUpdateToEntity(updApp, ent)

	if err = s.repository.Update(ctx, ent); errors.Is(err, gorm.ErrDuplicatedKey) {
		err = common.ErrApplicationWithSameNameAndVersionExists
		return
	}

	return s.FindById(ctx, ent.ID)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) (err error) {
	if err = s.validator.VarCtx(ctx, id, "uuid4_rfc4122"); err != nil {
		return
	}
	return s.repository.Delete(ctx, id)
}

func (s *Service) FindById(ctx context.Context, id uuid.UUID) (app *model.Application, err error) {
	if err = s.validator.VarCtx(ctx, id, "uuid4_rfc4122"); err != nil {
		return
	}

	var ent *entity.Application
	if ent, err = s.repository.FindById(ctx, id); err != nil {
		return
	}

	app = new(model.Application)
	s.converter.ToModel(ent, app)
	return
}
