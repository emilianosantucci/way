package service

import (
	"context"
	"libs/core/feature/resource/route/repository"
	"libs/core/feature/resource/route/repository/entity"
	"libs/core/feature/resource/route/service/model"

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

func (s *Service) Create(ctx context.Context, newRoute *model.NewRoute) (route *model.Route, err error) {
	if err = s.validator.StructCtx(ctx, newRoute); err != nil {
		return
	}

	ent := new(entity.Route)
	s.converter.FromNewToEntity(newRoute, ent)

	if err = s.repository.Create(ctx, ent); err != nil {
		return
	}

	route = new(model.Route)
	s.converter.ToModel(ent, route)

	return
}

func (s *Service) Update(ctx context.Context, updRoute *model.UpdateRoute) (route *model.Route, err error) {
	if err = s.validator.StructCtx(ctx, updRoute); err != nil {
		return
	}
	ent := new(entity.Route)
	s.converter.FromUpdateToEntity(updRoute, ent)

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

func (s *Service) FindById(ctx context.Context, id uuid.UUID) (route *model.Route, err error) {
	if err = s.validator.VarCtx(ctx, id, "uuid4_rfc4122"); err != nil {
		return
	}

	ent := new(entity.Route)
	if ent, err = s.repository.FindById(ctx, id); err != nil {
		return
	}

	route = new(model.Route)
	s.converter.ToModel(ent, route)

	return
}

func (s *Service) FindAll(ctx context.Context) (routes []*model.Route, err error) {
	ents := make([]*entity.Route, 0)
	if ents, err = s.repository.FindAll(ctx); err != nil {
		return
	}

	routes = make([]*model.Route, len(ents))
	for i, ent := range ents {
		routes[i] = new(model.Route)
		s.converter.ToModel(ent, routes[i])
	}

	return
}

func (s *Service) FindByPath(ctx context.Context, path string) (route *model.Route, err error) {
	if err = s.validator.VarCtx(ctx, path, "required,min=1"); err != nil {
		return
	}

	ent := new(entity.Route)
	if ent, err = s.repository.FindByPath(ctx, path); err != nil {
		return
	}

	route = new(model.Route)
	s.converter.ToModel(ent, route)

	return
}

func (s *Service) FindByName(ctx context.Context, name string) (route *model.Route, err error) {
	if err = s.validator.VarCtx(ctx, name, "required,min=1"); err != nil {
		return
	}

	ent := new(entity.Route)
	if ent, err = s.repository.FindByName(ctx, name); err != nil {
		return
	}

	route = new(model.Route)
	s.converter.ToModel(ent, route)

	return
}
