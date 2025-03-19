package usecase

import (
	"context"
	"schedulr-backend/services/user/internal/domain/dto"
	"schedulr-backend/services/user/internal/domain/entity"
	"schedulr-backend/services/user/internal/port/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"github.com/google/uuid"
	sharedError "github.com/virgiawanly/schedulr-backend/shared/errors"
)

type UserUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (u *UserUsecase) CreateUserFromRegistration(ctx context.Context, req dto.CreateUserFromRegistrationRequestDTO) (*dto.CreateUserFromRegistrationResponseDTO, *sharedError.AppError) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, sharedError.NewAppError(codes.Internal, "missing metadata")
	}

	businessIdHeaders := md.Get("business_id")
	if len(businessIdHeaders) <= 0 {
		return nil, sharedError.NewAppError(codes.InvalidArgument, "missing business id")
	}

	user := entity.User{
		ID:         uuid.New().String(),
		BusinessID: businessIdHeaders[0],
		AccountID:  req.AccountID,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Phone:      req.Phone,
		Address1:   req.Address1,
		Address2:   req.Address2,
		City:       req.City,
		State:      req.State,
		Zipcode:    req.Zipcode,
		Country:    req.Country,
	}

	err := u.userRepository.Create(ctx, &user, nil)
	if err != nil {
		return nil, sharedError.WrapError(codes.Internal, err)
	}

	return &dto.CreateUserFromRegistrationResponseDTO{
		User: user,
	}, nil
}
