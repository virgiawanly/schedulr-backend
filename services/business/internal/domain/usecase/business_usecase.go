package usecase

import (
	"context"
	"schedulr-backend/services/business/internal/domain/dto"
	"schedulr-backend/services/business/internal/domain/entity"
	"schedulr-backend/services/business/internal/port/repository"

	"google.golang.org/grpc/codes"

	"github.com/google/uuid"
	sharedError "github.com/virgiawanly/schedulr-backend/shared/errors"
)

type BusinessUsecase struct {
	businessRepository repository.BusinessRepository
}

func NewBusinessUsecase(businessRepository repository.BusinessRepository) *BusinessUsecase {
	return &BusinessUsecase{businessRepository: businessRepository}
}

func (u *BusinessUsecase) CreateBusinessFromRegistration(ctx context.Context, reqDto *dto.CreateBusinessFromRegistrationRequestDTO) (*dto.CreateBusinessFromRegistrationResponseDTO, *sharedError.AppError) {
	business := entity.Business{
		ID:       uuid.New().String(),
		Name:     reqDto.Name,
		Industry: reqDto.Industry,
		Address1: reqDto.Address1,
		Address2: reqDto.Address2,
		City:     reqDto.City,
		State:    reqDto.State,
		Zipcode:  reqDto.Zipcode,
		Country:  reqDto.Country,
	}

	err := u.businessRepository.Create(ctx, &business, nil)
	if err != nil {
		return nil, sharedError.WrapError(codes.Internal, err)
	}

	return &dto.CreateBusinessFromRegistrationResponseDTO{
		Business: &business,
	}, nil
}
