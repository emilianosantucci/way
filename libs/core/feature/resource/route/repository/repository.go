package repository

import (
	"context"
	"libs/core/common"
	"libs/core/feature/resource/route/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func RegisterEntities(db *gorm.DB) error {
	return db.AutoMigrate(&entity.Route{})
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) withTx(tx *gorm.DB) *Repository {
	return &Repository{db: tx}
}

func (r *Repository) Transaction(ctx context.Context, fn func(repo *Repository) error) (err error) {
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

func (r *Repository) Create(ctx context.Context, ent *entity.Route) (err error) {
	return gorm.G[entity.Route](r.db).Create(ctx, ent)
}

func (r *Repository) Update(ctx context.Context, ent *entity.Route) (err error) {
	var rowsAffected int
	rowsAffected, err = gorm.G[entity.Route](r.db).Where("id = ?", ent.ID).Updates(ctx, *ent)
	return common.GenerateEmptyRowsAffectedError(err, rowsAffected, common.ErrRouteNotFound)
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	var rowsAffected int
	rowsAffected, err = gorm.G[entity.Route](r.db).Where("id = ?", id).Delete(ctx)
	return common.GenerateEmptyRowsAffectedError(err, rowsAffected, common.ErrRouteNotFound)
}

func (r *Repository) FindById(ctx context.Context, id uuid.UUID) (ent *entity.Route, err error) {
	ent, err = gorm.G[*entity.Route](r.db).Where("id = ?", id).First(ctx)
	return ent, common.GenerateRecordNotFoundError(err, common.ErrRouteNotFound)
}

func (r *Repository) FindAll(ctx context.Context) (ents []*entity.Route, err error) {
	ents, err = gorm.G[*entity.Route](r.db).Find(ctx)
	return
}

func (r *Repository) FindByPath(ctx context.Context, path string) (ent *entity.Route, err error) {
	ent, err = gorm.G[*entity.Route](r.db).Where("path = ?", path).First(ctx)
	return ent, common.GenerateRecordNotFoundError(err, common.ErrRouteNotFound)
}

func (r *Repository) FindByName(ctx context.Context, name string) (ent *entity.Route, err error) {
	ent, err = gorm.G[*entity.Route](r.db).Where("name = ?", name).First(ctx)
	return ent, common.GenerateRecordNotFoundError(err, common.ErrRouteNotFound)
}
