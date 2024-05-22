package dto

type AddressDTO struct {
	Street     string `json:"street" validate:"required"`
	Number     string `json:"number" validate:"required"`
	Complement string `json:"complement" validate:"required"`
	District   string `json:"district" validate:"required"`
	City       string `json:"city" validate:"required"`
	State      string `json:"state" validate:"required"`
	ZipCode    string `json:"zip_code" validate:"required"`
}
