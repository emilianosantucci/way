package repository

import (
	"context"
	"libs/core/common"
	"libs/core/feature/resource/restapi/repository/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func RegisterEntities(db *gorm.DB) error {
	return db.AutoMigrate(&entity.RestApiResource{})
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

func (r *Repository) Create(ctx context.Context, ent *entity.RestApiResource) (err error) {
	return gorm.G[entity.RestApiResource](r.db).Create(ctx, ent)
}

func (r *Repository) Update(ctx context.Context, ent *entity.RestApiResource) (err error) {
	var rowsAffected int
	rowsAffected, err = gorm.G[entity.RestApiResource](r.db).Where("id = ?", ent.ID).Updates(ctx, *ent)
	if err == nil && rowsAffected == 0 {
		err = common.ErrRestApiResourceNotFound
	}
	return
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	var rowsAffected int
	rowsAffected, err = gorm.G[entity.RestApiResource](r.db).Where("id = ?", id).Delete(ctx)
	if err == nil && rowsAffected == 0 {
		err = common.ErrRestApiResourceNotFound
	}
	return
}

func (r *Repository) FindById(ctx context.Context, id uuid.UUID) (ent *entity.RestApiResource, err error) {
	return gorm.G[*entity.RestApiResource](r.db).Where("id = ?", id).First(ctx)
}
