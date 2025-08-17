package repository

import (
	"context"

	"libs/core/application/model"
	"libs/core/common"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) withTx(tx *gorm.DB) *Repository {
	return &Repository{db: tx}
}

func (r *Repository) Transaction(ctx context.Context, fn func(repo *Repository) error) (err error) {
	tx := r.db.WithContext(ctx).Begin()

	if tx.Error != nil {
		return tx.Error
	}

	repo := r.withTx(tx)
	err = fn(repo)

	if err != nil {
		tx.Rollback()
		return
	}

	return tx.Commit().Error
}

func (r *Repository) Create(ctx context.Context, entity *model.Application) (err error) {
	return gorm.G[model.Application](r.db).Create(ctx, entity)
}

func (r *Repository) Update(ctx context.Context, entity *model.Application) (err error) {
	var rowsAffected int
	rowsAffected, err = gorm.G[model.Application](r.db).Where("id = ?", entity.ID).Updates(ctx, *entity)
	if err == nil && rowsAffected == 0 {
		err = common.ErrApplicationNotFound
	}
	return
}

func (r *Repository) FindById(ctx context.Context, id uuid.UUID) (entity *model.Application, err error) {
	return gorm.G[*model.Application](r.db).Where("id = ?", id).First(ctx)
}

func (r *Repository) FindByNameAndVersion(ctx context.Context, name string, version string) (entity *model.Application, err error) {
	return gorm.G[*model.Application](r.db).Where("name = ? AND version = ?", name, version).First(ctx)
}
