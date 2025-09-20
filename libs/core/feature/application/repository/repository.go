package repository

import (
	"context"
	"libs/core/common"
	"libs/core/feature/application/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func RegisterEntities(db *gorm.DB) error {
	return db.AutoMigrate(&entity.Application{})
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

func (r *Repository) Create(ctx context.Context, ent *entity.Application) (err error) {
	return gorm.G[entity.Application](r.db).Create(ctx, ent)
}

func (r *Repository) Update(ctx context.Context, ent *entity.Application) (err error) {
	var rowsAffected int
	rowsAffected, err = gorm.G[entity.Application](r.db).Where("id = ?", ent.ID).Updates(ctx, *ent)
	return common.GenerateEmptyRowsAffectedError(err, rowsAffected, common.ErrRestApiResourceNotFound)
}

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	var rowsAffected int
	rowsAffected, err = gorm.G[entity.Application](r.db).Where("id = ?", id).Delete(ctx)
	return common.GenerateEmptyRowsAffectedError(err, rowsAffected, common.ErrRestApiResourceNotFound)
}

func (r *Repository) FindById(ctx context.Context, id uuid.UUID) (ent *entity.Application, err error) {
	ent, err = gorm.G[*entity.Application](r.db).Where("id = ?", id).First(ctx)
	return ent, common.GenerateRecordNotFoundError(err, common.ErrApplicationNotFound)
}

func (r *Repository) FindByNameAndVersion(ctx context.Context, name string, version string) (ent *entity.Application, err error) {
	ent, err = gorm.G[*entity.Application](r.db).Where("name = ? AND version = ?", name, version).First(ctx)
	return ent, common.GenerateRecordNotFoundError(err, common.ErrApplicationNotFound)
}

func (r *Repository) FindAll(ctx context.Context) (apps []entity.Application, err error) {
	return gorm.G[entity.Application](r.db).Find(ctx)
}
