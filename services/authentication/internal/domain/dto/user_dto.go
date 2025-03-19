package dto

type CreateUserFromRegistrationRequestDTO struct {
	AccountID string  `json:"account_id"`
	FirstName string  `json:"first_name"`
	LastName  *string `json:"last_name"`
	Phone     *string `json:"phone"`
	Address1  *string `json:"address_1"`
	Address2  *string `json:"address_2"`
	City      *string `json:"city"`
	State     *string `json:"state"`
	Zipcode   *string `json:"zipcode"`
	Country   *string `json:"country"`
}

type CreateUserFromRegistrationResponseDTO struct {
	User any `json:"user"`
}
