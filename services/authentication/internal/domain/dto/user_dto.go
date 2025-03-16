package dto

type FindUserByIdRequestDTO struct {
	ID string `json:"id"`
}

type FindUserByIdResponseDTO struct {
	User any `json:"user"`
}

type FindUserByAccountIdRequestDTO struct {
	AccountID string `json:"account_id"`
}

type FindUserByAccountIdResponseDTO struct {
	UserID any `json:"user"`
}

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

type UpdateUserByAccountIdRequestDTO struct {
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

type UpdateUserByAccountIdResponseDTO struct {
	User any `json:"user"`
}
