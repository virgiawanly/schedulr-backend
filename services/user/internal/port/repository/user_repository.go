package repository

import (
	"context"
	"schedulr-backend/services/user/internal/domain/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	BeginTransaction() *gorm.DB
	Create(ctx context.Context, user *entity.User, tx *gorm.DB) error
}
