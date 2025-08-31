package repository

import (
	"context"
	"libs/core/common"
	"libs/core/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewRestApiResourceRepository(db *gorm.DB) *RestApiResourceRepository {
	return &RestApiResourceRepository{db: db}
}

type RestApiResourceRepository struct {
	db *gorm.DB
}

func (r *RestApiResourceRepository) withTx(tx *gorm.DB) *RestApiResourceRepository {
	return &RestApiResourceRepository{db: tx}
}

func (r *RestApiResourceRepository) Transaction(ctx context.Context, fn func(repo *RestApiResourceRepository) error) (err error) {
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

func (r *RestApiResourceRepository) Create(ctx context.Context, ent *entity.RestApiResource) (err error) {
	return gorm.G[entity.RestApiResource](r.db).Create(ctx, ent)
}

func (r *RestApiResourceRepository) Update(ctx context.Context, ent *entity.RestApiResource) (err error) {
	var rowsAffected int
	rowsAffected, err = gorm.G[entity.RestApiResource](r.db).Where("id = ?", ent.ID).Updates(ctx, *ent)
	return common.GenerateEmptyRowsAffectedError(err, rowsAffected, common.ErrRestApiResourceNotFound)
}

func (r *RestApiResourceRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	var rowsAffected int
	rowsAffected, err = gorm.G[entity.RestApiResource](r.db).Where("id = ?", id).Delete(ctx)
	return common.GenerateEmptyRowsAffectedError(err, rowsAffected, common.ErrRestApiResourceNotFound)
}

func (r *RestApiResourceRepository) FindById(ctx context.Context, id uuid.UUID) (ent *entity.RestApiResource, err error) {
	ent, err = gorm.G[*entity.RestApiResource](r.db).Where("id = ?", id).First(ctx)
	return ent, common.GenerateRecordNotFoundError(err, common.ErrRestApiResourceNotFound)
}
