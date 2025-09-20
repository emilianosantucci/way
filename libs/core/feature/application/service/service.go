package service

import (
	"context"
	"errors"
	"libs/core/common"
	"libs/core/feature/application/entity"
	"libs/core/feature/application/mapper"
	"libs/core/feature/application/model"
	"libs/core/feature/application/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewService(repository *repository.Repository, validator *validator.Validate, mapper mapper.ModelMap) *Service {
	return &Service{
		repository: repository,
		validator:  validator,
		mapper:     mapper,
	}
}

type Service struct {
	repository *repository.Repository
	validator  *validator.Validate
	mapper     mapper.ModelMap
}

func (s *Service) Create(ctx context.Context, newApp *model.NewApplication) (app *model.Application, err error) {
	if err = s.validator.StructCtx(ctx, newApp); err != nil {
		return
	}

	ent := new(entity.Application)
	s.mapper.FromNewToEntity(newApp, ent)

	if err = s.repository.Create(ctx, ent); err != nil {
		return
	}

	app = new(model.Application)
	s.mapper.ToModel(ent, app)

	return
}

func (s *Service) Update(ctx context.Context, updApp *model.UpdateApplication) (app *model.Application, err error) {
	if err = s.validator.StructCtx(ctx, updApp); err != nil {
		return
	}

	ent := new(entity.Application)
	s.mapper.FromUpdateToEntity(updApp, ent)

	if err = s.repository.Update(ctx, ent); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			err = common.ErrApplicationWithSameNameAndVersionExists
		}
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
	s.mapper.ToModel(ent, app)
	return
}

func (s *Service) FindAll(ctx context.Context) (apps []model.Application, err error) {
	var entities []entity.Application
	entities, err = s.repository.FindAll(ctx)

	if err != nil {
		return
	}

	apps = s.mapper.ToModels(entities)
	return
}
