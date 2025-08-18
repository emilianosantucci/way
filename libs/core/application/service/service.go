package service

import (
	"context"
	"errors"
	"libs/core/application/repository"
	"libs/core/application/repository/entity"
	"libs/core/application/service/model"
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

	ent := new(entity.Application)
	if err = copier.Copy(ent, newApp); err != nil {
		return
	}

	if err = s.repository.Create(ctx, ent); err != nil {
		return
	}

	return s.transformEntityToModel(ent)
}

func (s *Service) Update(ctx context.Context, updApp *model.UpdateApplication) (app *model.Application, err error) {
	if err = s.validator.StructCtx(ctx, updApp); err != nil {
		return
	}

	ent := new(entity.Application)
	if err = copier.Copy(ent, updApp); err != nil {
		return
	}

	if err = s.repository.Update(ctx, ent); errors.Is(err, gorm.ErrDuplicatedKey) {
		err = common.ErrApplicationWithSameNameAndVersionExists
	}

	return s.transformEntityToModel(ent)
}

func (s *Service) FindById(ctx context.Context, id uuid.UUID) (app *model.Application, err error) {
	if err = s.validator.VarCtx(ctx, id, "uuid4_rfc4122"); err != nil {
		return
	}

	var ent *entity.Application
	if ent, err = s.repository.FindById(ctx, id); err != nil {
		return
	}

	return s.transformEntityToModel(ent)
}

func (s *Service) transformEntityToModel(ent *entity.Application) (mod *model.Application, err error) {
	mod = new(model.Application)
	return mod, copier.Copy(mod, ent)
}
