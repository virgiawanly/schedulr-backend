package dto

import "schedulr-backend/services/business/internal/domain/entity"

type CreateBusinessFromRegistrationRequestDTO struct {
	Name        string  `json:"name" validate:"required,max=255"`
	Industry    string  `json:"industry" validate:"required,max=255"`
	CompanySize *string `json:"company_size,omitempty" validate:"omitempty,max=255"`
	Address1    *string `json:"address_1,omitempty" validate:"omitempty,max=255"`
	Address2    *string `json:"address_2,omitempty" validate:"omitempty,max=255"`
	City        *string `json:"city,omitempty" validate:"omitempty,max=150"`
	State       *string `json:"state,omitempty" validate:"omitempty,max=150"`
	Zipcode     *string `json:"zipcode,omitempty" validate:"omitempty,max=20"`
	Country     *string `json:"country,omitempty" validate:"omitempty,max=100"`
}

type CreateBusinessFromRegistrationResponseDTO struct {
	Business *entity.Business `json:"business"`
}
