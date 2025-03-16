package mapper

import (
	"schedulr-backend/services/user/internal/domain/dto"

	userV1 "github.com/virgiawanly/schedulr-backend/protogen/go/user/v1"
)

func ToCreateUserFromRegistrationRequestDTO(req *userV1.CreateUserFromRegistrationRequest) dto.CreateUserFromRegistrationRequestDTO {
	return dto.CreateUserFromRegistrationRequestDTO{
		AccountID: req.AccountId,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Phone:     req.Phone,
		Address1:  req.Address_1,
		Address2:  req.Address_2,
		City:      req.City,
		State:     req.State,
		Zipcode:   req.Zipcode,
		Country:   req.Country,
	}
}
