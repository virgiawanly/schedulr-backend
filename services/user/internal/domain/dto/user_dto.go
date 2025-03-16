package dto

import "schedulr-backend/services/user/internal/domain/entity"

type CreateUserFromRegistrationRequestDTO struct {
	AccountID string  `json:"account_id" validate:"omitempty"`
	FirstName string  `json:"first_name" validate:"omitempty,max=255"`
	LastName  *string `json:"last_name" validate:"omitempty,max=255"`
	Email     string  `json:"email" validate:"omitempty,max=255"`
	Password  string  `json:"password" validate:"omitempty,min=5"`
	Phone     *string `json:"phone" validate:"omitempty,max=25"`
	Address1  *string `json:"address_1" validate:"omitempty,max=255"`
	Address2  *string `json:"address_2" validate:"omitempty,max=255"`
	City      *string `json:"city" validate:"omitempty,max=150"`
	State     *string `json:"state" validate:"omitempty,max=150"`
	Zipcode   *string `json:"zipcode" validate:"omitempty,max=20"`
	Country   *string `json:"country" validate:"omitempty,max=100"`
}

type CreateUserFromRegistrationResponseDTO struct {
	User entity.User `json:"user"`
}
