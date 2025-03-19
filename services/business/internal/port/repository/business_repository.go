package repository

import (
	"context"
	"schedulr-backend/services/business/internal/domain/entity"

	"gorm.io/gorm"
)

type BusinessRepository interface {
	BeginTransaction() *gorm.DB
	Create(ctx context.Context, business *entity.Business, tx *gorm.DB) error
}
