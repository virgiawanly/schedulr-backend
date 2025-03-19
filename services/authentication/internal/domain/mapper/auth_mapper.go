package mapper

import (
	"fmt"
	"schedulr-backend/services/authentication/internal/domain/dto"
	"schedulr-backend/services/authentication/internal/domain/entity"

	authenticationV1 "github.com/virgiawanly/schedulr-backend/protogen/go/authentication/v1"
)

func ToRegisterRequestDTO(req *authenticationV1.RegisterRequest) dto.RegisterRequestDTO {
	return dto.RegisterRequestDTO{
		FirstName:           req.FirstName,
		LastName:            req.LastName,
		Email:               req.Email,
		Password:            req.Password,
		Phone:               req.Phone,
		Address1:            req.Address_1,
		Address2:            req.Address_2,
		City:                req.City,
		State:               req.State,
		Zipcode:             req.Zipcode,
		Country:             req.Country,
		BusinessName:        req.BusinessName,
		BusinessIndustry:    req.BusinessIndustry,
		BusinessCompanySize: req.BusinessCompanySize,
	}
}

func ToRegisterResponseDTO(account entity.Account) dto.RegisterResponseDTO {
	return dto.RegisterResponseDTO{
		AccountID: account.ID,
	}
}

func ToCreateBusinessFromRegistrationRequestDTO(req dto.RegisterRequestDTO) dto.CreateBusinessFromRegistrationRequestDTO {
	var businessName string
	if req.BusinessName != nil {
		businessName = *req.BusinessName
	} else {
		businessName = fmt.Sprintf("%v %v", req.FirstName, req.LastName)
	}

	return dto.CreateBusinessFromRegistrationRequestDTO{
		Name:        businessName,
		Industry:    req.BusinessIndustry,
		CompanySize: req.BusinessCompanySize,
		Address1:    req.Address1,
		Address2:    req.Address2,
		City:        req.City,
		State:       req.State,
		Zipcode:     req.Zipcode,
		Country:     req.Country,
	}
}

func ToCreateUserFromRegistrationRequestDTO(req dto.RegisterRequestDTO, accountId string) dto.CreateUserFromRegistrationRequestDTO {
	return dto.CreateUserFromRegistrationRequestDTO{
		AccountID: accountId,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Address1:  req.Address1,
		Address2:  req.Address2,
		City:      req.City,
		State:     req.State,
		Zipcode:   req.Zipcode,
		Country:   req.Country,
	}
}
