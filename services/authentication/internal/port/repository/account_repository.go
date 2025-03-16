package repository

import (
	"context"
	"schedulr-backend/services/authentication/internal/domain/entity"

	"gorm.io/gorm"
)

type AccountRepository interface {
	BeginTransaction() *gorm.DB
	Create(ctx context.Context, account *entity.Account, tx *gorm.DB) error
	FindByEmail(ctx context.Context, email string) (*entity.Account, error)
	CheckEmailExists(ctx context.Context, email string) (bool, error)
}
