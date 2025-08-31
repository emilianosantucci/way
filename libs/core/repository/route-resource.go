package repository

import (
	"context"
	"libs/core/common"
	"libs/core/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewRouteResourceRepository(db *gorm.DB) *RouteResourceRepository {
	return &RouteResourceRepository{db: db}
}

type RouteResourceRepository struct {
	db *gorm.DB
}

func (r *RouteResourceRepository) withTx(tx *gorm.DB) *RouteResourceRepository {
	return &RouteResourceRepository{db: tx}
}

func (r *RouteResourceRepository) Transaction(ctx context.Context, fn func(repo *RouteResourceRepository) error) (err error) {
	var tx *gorm.DB
	if tx = r.db.WithContext(ctx).Begin(); tx.Error != nil {
		return tx.Error
	}

	repo := r.withTx(tx)

	if err = fn(repo); err != nil {
		tx.Rollback()
		return
	}

	return tx.Commit().Error
}

func (r *RouteResourceRepository) Create(ctx context.Context, ent *entity.Route) (err error) {
	return gorm.G[entity.Route](r.db).Create(ctx, ent)
}

func (r *RouteResourceRepository) Update(ctx context.Context, ent *entity.Route) (err error) {
	var rowsAffected int
	rowsAffected, err = gorm.G[entity.Route](r.db).Where("id = ?", ent.ID).Updates(ctx, *ent)
	return common.GenerateEmptyRowsAffectedError(err, rowsAffected, common.ErrRouteNotFound)
}

func (r *RouteResourceRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	var rowsAffected int
	rowsAffected, err = gorm.G[entity.Route](r.db).Where("id = ?", id).Delete(ctx)
	return common.GenerateEmptyRowsAffectedError(err, rowsAffected, common.ErrRouteNotFound)
}

func (r *RouteResourceRepository) FindById(ctx context.Context, id uuid.UUID) (ent *entity.Route, err error) {
	ent, err = gorm.G[*entity.Route](r.db).Where("id = ?", id).First(ctx)
	return ent, common.GenerateRecordNotFoundError(err, common.ErrRouteNotFound)
}

func (r *RouteResourceRepository) FindAll(ctx context.Context) (ents []*entity.Route, err error) {
	ents, err = gorm.G[*entity.Route](r.db).Find(ctx)
	return
}

func (r *RouteResourceRepository) FindByPath(ctx context.Context, path string) (ent *entity.Route, err error) {
	ent, err = gorm.G[*entity.Route](r.db).Where("path = ?", path).First(ctx)
	return ent, common.GenerateRecordNotFoundError(err, common.ErrRouteNotFound)
}

func (r *RouteResourceRepository) FindByName(ctx context.Context, name string) (ent *entity.Route, err error) {
	ent, err = gorm.G[*entity.Route](r.db).Where("name = ?", name).First(ctx)
	return ent, common.GenerateRecordNotFoundError(err, common.ErrRouteNotFound)
}
