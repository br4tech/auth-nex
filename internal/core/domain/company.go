package domain

type Company struct {
	LegalName         string `validate:"required"`
	TradeName         string `validate:"required"`
	CNPJ              string `validate:"required"`
	StateRegistration string `validate:"required"`
	Address           Address
	Partners          []Partner
	Activities        []Activity
	Type              string // Company type (MEI, ME, LTDA, etc.)
	TenantID          uint   `validate:"required"`
	Schema            string
	Users             []User
}

func NewCompany(legalName string, tradeName string, cnpj string,
	stateRegistration string, address Address, partners []Partner,
	activities []Activity, ctype string, tenantId uint, schema string,
	users []User) *Company {
	return &Company{
		LegalName:         legalName,
		TradeName:         tradeName,
		CNPJ:              cnpj,
		StateRegistration: stateRegistration,
		Address:           address,
		Partners:          partners,
		Activities:        activities,
		Type:              ctype,
		TenantID:          tenantId,
		Schema:            schema,
		Users:             users,
	}
}
