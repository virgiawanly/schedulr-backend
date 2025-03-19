package grpc_client

import (
	"context"
	"schedulr-backend/services/authentication/internal/domain/dto"
)

type BusinessClientPort interface {
	CreateBusinessFromRegistration(ctx context.Context, dto *dto.CreateBusinessFromRegistrationRequestDTO) (*dto.CreateBusinessFromRegistrationResponseDTO, error)
}
