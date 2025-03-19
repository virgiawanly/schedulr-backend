package dto

import "time"

type CheckRegisteredEmailRequestDTO struct {
	Email string `json:"email" validate:"required,email"`
}

type RegisterRequestDTO struct {
	FirstName           string  `json:"first_name" validate:"required,max=255"`
	LastName            *string `json:"last_name" validate:"omitempty,max=255"`
	Email               string  `json:"email" validate:"required,email,max=255"`
	Password            string  `json:"password" validate:"required,min=5"`
	Phone               *string `json:"phone" validate:"omitempty,max=25"`
	Address1            *string `json:"address_1" validate:"omitempty,max=255"`
	Address2            *string `json:"address_2" validate:"omitempty,max=255"`
	City                *string `json:"city" validate:"omitempty,max=150"`
	State               *string `json:"state" validate:"omitempty,max=150"`
	Zipcode             *string `json:"zipcode" validate:"omitempty,max=20"`
	Country             *string `json:"country" validate:"omitempty,max=100"`
	BusinessName        *string `json:"business_name" validate:"omitempty,max=255"`
	BusinessIndustry    string  `json:"business_industry" validate:"required,max=255"`
	BusinessCompanySize *string `json:"business_company_size" validate:"omitempty,max=20"`
}

type RegisterResponseDTO struct {
	AccountID string `json:"account_id"`
}

type LoginRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	TokenType    string        `json:"token_type"`
	ExpiryTime   time.Duration `json:"expiry_time"`
}
