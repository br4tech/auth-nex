package dto

type CompanyDTO struct {
	LegalName         string     `json:"legal_name" validate:"required"`
	TradeName         string     `json:"trade_name" validate:"required"`
	Document          string     `json:"document" validate:"required"`
	StateRegistration string     `json:"state_registration" validate:"required"`
	Address           AddressDTO `json:"adress" validate:"required"`
	Type              string     `json:"type" validate:"required"`
}
