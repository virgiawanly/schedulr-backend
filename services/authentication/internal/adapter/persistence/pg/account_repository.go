package pg

import (
	"context"
	"schedulr-backend/services/authentication/internal/domain/entity"
	"schedulr-backend/services/authentication/internal/port/repository"

	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) repository.AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

func (r *accountRepository) Create(ctx context.Context, account *entity.Account, tx *gorm.DB) error {
	if tx != nil {
		return tx.WithContext(ctx).Create(account).Error
	}
	return r.db.WithContext(ctx).Create(account).Error
}

func (r *accountRepository) FindByEmail(ctx context.Context, email string) (*entity.Account, error) {
	var account entity.Account
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&entity.Account{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
