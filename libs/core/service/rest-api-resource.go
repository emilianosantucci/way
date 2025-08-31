package service

import (
	"context"
	"libs/core/entity"
	"libs/core/model"
	"libs/core/model/convert"
	"libs/core/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewRestApiResourceService(repository *repository.RestApiResourceRepository, validator *validator.Validate, converter convert.RestApiResourceConvert) (service *RestApiResourceService) {
	return &RestApiResourceService{
		repository: repository,
		validator:  validator,
		converter:  converter,
	}
}

type RestApiResourceService struct {
	repository *repository.RestApiResourceRepository
	validator  *validator.Validate
	converter  convert.RestApiResourceConvert
}

func (s *RestApiResourceService) Create(ctx context.Context, newRest *model.NewRestApiResource) (rest *model.RestApiResource, err error) {
	if err = s.validator.StructCtx(ctx, newRest); err != nil {
		return
	}

	ent := new(entity.RestApiResource)
	s.converter.FromNewToEntity(newRest, ent)

	if err = s.repository.Create(ctx, ent); err != nil {
		return
	}

	rest = new(model.RestApiResource)
	s.converter.ToModel(ent, rest)

	return
}

func (s *RestApiResourceService) Update(ctx context.Context, updRest *model.UpdateRestApiResource) (rest *model.RestApiResource, err error) {
	if err = s.validator.StructCtx(ctx, updRest); err != nil {
		return
	}
	ent := new(entity.RestApiResource)
	s.converter.FromUpdateToEntity(updRest, ent)

	if err = s.repository.Update(ctx, ent); err != nil {
		return
	}

	return s.FindById(ctx, ent.ID)
}

func (s *RestApiResourceService) Delete(ctx context.Context, id uuid.UUID) (err error) {
	if err = s.validator.VarCtx(ctx, id, "uuid4_rfc4122"); err != nil {
		return
	}

	return s.repository.Delete(ctx, id)
}

func (s *RestApiResourceService) FindById(ctx context.Context, id uuid.UUID) (rest *model.RestApiResource, err error) {
	if err = s.validator.VarCtx(ctx, id, "uuid4_rfc4122"); err != nil {
		return
	}

	ent := new(entity.RestApiResource)
	if ent, err = s.repository.FindById(ctx, id); err != nil {
		return
	}

	rest = new(model.RestApiResource)
	s.converter.ToModel(ent, rest)

	return
}
