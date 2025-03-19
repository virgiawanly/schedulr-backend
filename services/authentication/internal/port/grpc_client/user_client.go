package grpc_client

import (
	"context"
	"schedulr-backend/services/authentication/internal/domain/dto"
)

type UserClientPort interface {
	CreateUserFromRegistration(ctx context.Context, businessId *string, dto *dto.CreateUserFromRegistrationRequestDTO) (*dto.CreateUserFromRegistrationResponseDTO, error)
}
