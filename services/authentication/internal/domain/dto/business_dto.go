package dto

type BusinessResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Industry string `json:"industry"`
}

type CreateBusinessFromRegistrationRequestDTO struct {
	Name        string  `json:"name"`
	Industry    string  `json:"industry"`
	CompanySize *string `json:"company_size"`
	Address1    *string `json:"address_1"`
	Address2    *string `json:"address_2"`
	City        *string `json:"city"`
	State       *string `json:"state"`
	Zipcode     *string `json:"zipcode"`
	Country     *string `json:"country"`
}

type CreateBusinessFromRegistrationResponseDTO struct {
	Business BusinessResponse `json:"business"`
}
