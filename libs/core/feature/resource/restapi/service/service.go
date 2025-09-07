package service

import (
	"context"
	"libs/core/feature/resource/restapi/entity"
	"libs/core/feature/resource/restapi/mapper"
	"libs/core/feature/resource/restapi/model"
	"libs/core/feature/resource/restapi/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewService(repository *repository.Repository, validator *validator.Validate, mapper mapper.ModelMap) (service *Service) {
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

func (s *Service) Create(ctx context.Context, newRest *model.NewRestApiResource) (rest *model.RestApiResource, err error) {
	if err = s.validator.StructCtx(ctx, newRest); err != nil {
		return
	}

	ent := new(entity.RestApiResource)
	s.mapper.FromNewToEntity(newRest, ent)

	if err = s.repository.Create(ctx, ent); err != nil {
		return
	}

	rest = new(model.RestApiResource)
	s.mapper.ToModel(ent, rest)

	return
}

func (s *Service) Update(ctx context.Context, updRest *model.UpdateRestApiResource) (rest *model.RestApiResource, err error) {
	if err = s.validator.StructCtx(ctx, updRest); err != nil {
		return
	}
	ent := new(entity.RestApiResource)
	s.mapper.FromUpdateToEntity(updRest, ent)

	if err = s.repository.Update(ctx, ent); err != nil {
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

func (s *Service) FindById(ctx context.Context, id uuid.UUID) (rest *model.RestApiResource, err error) {
	if err = s.validator.VarCtx(ctx, id, "uuid4_rfc4122"); err != nil {
		return
	}

	ent := new(entity.RestApiResource)
	if ent, err = s.repository.FindById(ctx, id); err != nil {
		return
	}

	rest = new(model.RestApiResource)
	s.mapper.ToModel(ent, rest)

	return
}
