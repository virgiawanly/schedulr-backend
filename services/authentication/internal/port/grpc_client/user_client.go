package grpc_client

import (
	"context"
	"schedulr-backend/services/authentication/internal/domain/dto"
)

type UserClientPort interface {
	FindUserById(ctx context.Context, businessId *string, dto *dto.FindUserByIdRequestDTO) (*dto.FindUserByIdResponseDTO, error)
	CreateUserFromRegistration(ctx context.Context, businessId *string, dto *dto.CreateUserFromRegistrationRequestDTO) (*dto.CreateUserFromRegistrationResponseDTO, error)
	UpdateUserByAccountId(ctx context.Context, businessId *string, dto *dto.UpdateUserByAccountIdRequestDTO) (*dto.UpdateUserByAccountIdResponseDTO, error)
}
