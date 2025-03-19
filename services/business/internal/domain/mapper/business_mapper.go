package mapper

import (
	"schedulr-backend/services/business/internal/domain/dto"

	businessV1 "github.com/virgiawanly/schedulr-backend/protogen/go/business/v1"
)

func ToCreateBusinessFromRegistrationRequestDTO(req *businessV1.CreateBusinessFromRegistrationRequest) dto.CreateBusinessFromRegistrationRequestDTO {
	return dto.CreateBusinessFromRegistrationRequestDTO{
		Name:     req.Name,
		Industry: req.Industry,
		Address1: req.Address_1,
		Address2: req.Address_2,
		City:     req.City,
		State:    req.State,
		Zipcode:  req.Zipcode,
		Country:  req.Country,
	}
}
