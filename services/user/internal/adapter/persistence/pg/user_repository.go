package pg

import (
	"context"
	"schedulr-backend/services/user/internal/domain/entity"
	"schedulr-backend/services/user/internal/port/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

func (r *userRepository) Create(ctx context.Context, user *entity.User, tx *gorm.DB) error {
	if tx != nil {
		return tx.WithContext(ctx).Create(user).Error
	}
	return r.db.WithContext(ctx).Create(user).Error
}
