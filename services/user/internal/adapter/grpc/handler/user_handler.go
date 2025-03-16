package handler

import (
	"context"
	"schedulr-backend/services/user/internal/domain/mapper"
	"schedulr-backend/services/user/internal/domain/usecase"

	"github.com/go-playground/validator/v10"

	userV1 "github.com/virgiawanly/schedulr-backend/protogen/go/user/v1"
	errorutil "github.com/virgiawanly/schedulr-backend/shared/util/errors"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
	validator   *validator.Validate
	userV1.UnimplementedUserServiceServer
}

func NewUserHandler(usecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: usecase,
		validator:   validator.New(),
	}
}

func (h *UserHandler) CreateUserFromRegistration(ctx context.Context, req *userV1.CreateUserFromRegistrationRequest) (*userV1.CreateUserFromRegistrationResponse, error) {
	registerDto := mapper.ToCreateUserFromRegistrationRequestDTO(req)
	if err := h.validator.Struct(registerDto); err != nil {
		return nil, errorutil.MapToGRPCError(err)
	}

	result, err := h.userUsecase.CreateUserFromRegistration(ctx, registerDto)
	if err != nil {
		return nil, errorutil.MapToGRPCError(err)
	}

	return &userV1.CreateUserFromRegistrationResponse{User: &userV1.User{
		Id:         result.User.ID,
		AccountId:  result.User.AccountID,
		BusinessId: result.User.BusinessID,
		FirstName:  result.User.FirstName,
		LastName:   result.User.LastName,
		Phone:      result.User.Phone,
		Address_1:  result.User.Address1,
		Address_2:  result.User.Address2,
		City:       result.User.City,
		State:      result.User.State,
		Zipcode:    result.User.Zipcode,
	}}, nil
}
