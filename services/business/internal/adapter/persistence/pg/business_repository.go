package pg

import (
	"context"
	"schedulr-backend/services/business/internal/domain/entity"
	"schedulr-backend/services/business/internal/port/repository"

	"gorm.io/gorm"
)

type businessRepository struct {
	db *gorm.DB
}

func NewBusinessRepository(db *gorm.DB) repository.BusinessRepository {
	return &businessRepository{db: db}
}

func (r *businessRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

func (r *businessRepository) Create(ctx context.Context, business *entity.Business, tx *gorm.DB) error {
	if tx != nil {
		return tx.WithContext(ctx).Create(business).Error
	}
	return r.db.WithContext(ctx).Create(business).Error
}
