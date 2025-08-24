package service

import (
	"context"
	"libs/core/feature/resource/rest/repository"
	"libs/core/feature/resource/rest/repository/entity"
	"libs/core/feature/resource/rest/service/model"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewService(repository *repository.Repository, validator *validator.Validate, converter model.Convert) (service *Service) {
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

func (s *Service) Create(ctx context.Context, newRest *model.NewRestResource) (rest *model.RestResource, err error) {
	if err = s.validator.StructCtx(ctx, newRest); err != nil {
		return
	}

	ent := new(entity.RestResource)
	s.converter.FromNewToEntity(newRest, ent)

	if err = s.repository.Create(ctx, ent); err != nil {
		return
	}

	rest = new(model.RestResource)
	s.converter.ToModel(ent, rest)

	return
}

func (s *Service) Update(ctx context.Context, updRest *model.UpdateRestResource) (rest *model.RestResource, err error) {
	if err = s.validator.StructCtx(ctx, updRest); err != nil {
		return
	}
	ent := new(entity.RestResource)
	s.converter.FromUpdateToEntity(updRest, ent)

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

func (s *Service) FindById(ctx context.Context, id uuid.UUID) (rest *model.RestResource, err error) {
	if err = s.validator.VarCtx(ctx, id, "uuid4_rfc4122"); err != nil {
		return
	}

	ent := new(entity.RestResource)
	if ent, err = s.repository.FindById(ctx, id); err != nil {
		return
	}

	rest = new(model.RestResource)
	s.converter.ToModel(ent, rest)

	return
}
