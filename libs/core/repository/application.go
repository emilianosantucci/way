package repository

import (
	"context"
	"libs/core/common"
	"libs/core/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewApplicationRepository(db *gorm.DB) *ApplicationRepository {
	return &ApplicationRepository{db: db}
}

type ApplicationRepository struct {
	db *gorm.DB
}

func (r *ApplicationRepository) withTx(tx *gorm.DB) *RestApiResourceRepository {
	return &RestApiResourceRepository{db: tx}
}

func (r *ApplicationRepository) Transaction(ctx context.Context, fn func(repo *RestApiResourceRepository) error) (err error) {
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

func (r *ApplicationRepository) Create(ctx context.Context, ent *entity.Application) (err error) {
	return gorm.G[entity.Application](r.db).Create(ctx, ent)
}

func (r *ApplicationRepository) Update(ctx context.Context, ent *entity.Application) (err error) {
	var rowsAffected int
	rowsAffected, err = gorm.G[entity.Application](r.db).Where("id = ?", ent.ID).Updates(ctx, *ent)
	return common.GenerateEmptyRowsAffectedError(err, rowsAffected, common.ErrRestApiResourceNotFound)
}

func (r *ApplicationRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	var rowsAffected int
	rowsAffected, err = gorm.G[entity.Application](r.db).Where("id = ?", id).Delete(ctx)
	return common.GenerateEmptyRowsAffectedError(err, rowsAffected, common.ErrRestApiResourceNotFound)
}

func (r *ApplicationRepository) FindById(ctx context.Context, id uuid.UUID) (ent *entity.Application, err error) {
	ent, err = gorm.G[*entity.Application](r.db).Where("id = ?", id).First(ctx)
	return ent, common.GenerateRecordNotFoundError(err, common.ErrApplicationNotFound)
}

func (r *ApplicationRepository) FindByNameAndVersion(ctx context.Context, name string, version string) (ent *entity.Application, err error) {
	ent, err = gorm.G[*entity.Application](r.db).Where("name = ? AND version = ?", name, version).First(ctx)
	return ent, common.GenerateRecordNotFoundError(err, common.ErrApplicationNotFound)
}
