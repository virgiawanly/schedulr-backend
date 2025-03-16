package mapper

import (
	"schedulr-backend/services/authentication/internal/domain/dto"
	"schedulr-backend/services/authentication/internal/domain/entity"

	authenticationV1 "github.com/virgiawanly/schedulr-backend/protogen/go/authentication/v1"
)

func ToRegisterRequestDTO(req *authenticationV1.RegisterRequest) dto.RegisterRequestDTO {
	return dto.RegisterRequestDTO{
		FirstName: req.FirstName,
		LastName:  &req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Phone:     &req.Phone,
		Address1:  &req.Address_1,
		Address2:  &req.Address_2,
		City:      &req.City,
		State:     &req.State,
		Zipcode:   &req.Zipcode,
		Country:   &req.Country,
	}
}

func ToRegisterResponseDTO(account entity.Account) dto.RegisterResponseDTO {
	return dto.RegisterResponseDTO{
		AccountID: account.ID,
	}
}
