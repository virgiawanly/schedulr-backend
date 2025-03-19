package handler

import (
	"context"
	"schedulr-backend/services/business/internal/domain/mapper"
	"schedulr-backend/services/business/internal/domain/usecase"

	"github.com/go-playground/validator/v10"

	businessV1 "github.com/virgiawanly/schedulr-backend/protogen/go/business/v1"
	errorutil "github.com/virgiawanly/schedulr-backend/shared/util/errors"
)

type BusinessHandler struct {
	businessUsecase *usecase.BusinessUsecase
	validator       *validator.Validate
	businessV1.UnimplementedBusinessServiceServer
}

func NewBusinessHandler(usecase *usecase.BusinessUsecase) *BusinessHandler {
	return &BusinessHandler{
		businessUsecase: usecase,
		validator:       validator.New(),
	}
}

func (h *BusinessHandler) CreateBusinessFromRegistration(ctx context.Context, req *businessV1.CreateBusinessFromRegistrationRequest) (*businessV1.CreateBusinessFromRegistrationResponse, error) {
	registerDto := mapper.ToCreateBusinessFromRegistrationRequestDTO(req)
	if err := h.validator.Struct(registerDto); err != nil {
		return nil, errorutil.MapToGRPCError(err)
	}

	result, err := h.businessUsecase.CreateBusinessFromRegistration(ctx, &registerDto)
	if err != nil {
		return nil, errorutil.MapToGRPCError(err)
	}

	return &businessV1.CreateBusinessFromRegistrationResponse{Business: &businessV1.Business{
		Id:        result.Business.ID,
		Name:      result.Business.Name,
		Industry:  result.Business.Industry,
		Address_1: result.Business.Address1,
		Address_2: result.Business.Address2,
		City:      result.Business.City,
		State:     result.Business.State,
		Zipcode:   result.Business.Zipcode,
		Country:   result.Business.Country,
	}}, nil
}
