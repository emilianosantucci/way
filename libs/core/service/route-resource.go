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

func NewRouteResourceService(repository *repository.RouteResourceRepository, validator *validator.Validate, converter convert.RouteResourceConvert) (service *RouteResourceService) {
	return &RouteResourceService{
		repository: repository,
		validator:  validator,
		converter:  converter,
	}
}

type RouteResourceService struct {
	repository *repository.RouteResourceRepository
	validator  *validator.Validate
	converter  convert.RouteResourceConvert
}

func (s *RouteResourceService) Create(ctx context.Context, newRoute *model.NewRouteResource) (route *model.RouteResource, err error) {
	if err = s.validator.StructCtx(ctx, newRoute); err != nil {
		return
	}

	ent := new(entity.Route)
	s.converter.FromNewToEntity(newRoute, ent)

	if err = s.repository.Create(ctx, ent); err != nil {
		return
	}

	route = new(model.RouteResource)
	s.converter.ToModel(ent, route)

	return
}

func (s *RouteResourceService) Update(ctx context.Context, updRoute *model.UpdateRouteResource) (route *model.RouteResource, err error) {
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

func (s *RouteResourceService) Delete(ctx context.Context, id uuid.UUID) (err error) {
	if err = s.validator.VarCtx(ctx, id, "uuid4_rfc4122"); err != nil {
		return
	}

	return s.repository.Delete(ctx, id)
}

func (s *RouteResourceService) FindById(ctx context.Context, id uuid.UUID) (route *model.RouteResource, err error) {
	if err = s.validator.VarCtx(ctx, id, "uuid4_rfc4122"); err != nil {
		return
	}

	ent := new(entity.Route)
	if ent, err = s.repository.FindById(ctx, id); err != nil {
		return
	}

	route = new(model.RouteResource)
	s.converter.ToModel(ent, route)

	return
}

func (s *RouteResourceService) FindAll(ctx context.Context) (routes []*model.RouteResource, err error) {
	ents := make([]*entity.Route, 0)
	if ents, err = s.repository.FindAll(ctx); err != nil {
		return
	}

	routes = make([]*model.RouteResource, len(ents))
	for i, ent := range ents {
		routes[i] = new(model.RouteResource)
		s.converter.ToModel(ent, routes[i])
	}

	return
}

func (s *RouteResourceService) FindByPath(ctx context.Context, path string) (route *model.RouteResource, err error) {
	if err = s.validator.VarCtx(ctx, path, "required,min=1"); err != nil {
		return
	}

	ent := new(entity.Route)
	if ent, err = s.repository.FindByPath(ctx, path); err != nil {
		return
	}

	route = new(model.RouteResource)
	s.converter.ToModel(ent, route)

	return
}

func (s *RouteResourceService) FindByName(ctx context.Context, name string) (route *model.RouteResource, err error) {
	if err = s.validator.VarCtx(ctx, name, "required,min=1"); err != nil {
		return
	}

	ent := new(entity.Route)
	if ent, err = s.repository.FindByName(ctx, name); err != nil {
		return
	}

	route = new(model.RouteResource)
	s.converter.ToModel(ent, route)

	return
}
