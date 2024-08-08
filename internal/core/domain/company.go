package domain

type Company struct {
	Id                int
	LegalName         string `validate:"required"`
	TradeName         string `validate:"required"`
	Document          string `validate:"required"`
	StateRegistration string `validate:"required"`
	Address           Address
	Partners          []Partner
	Activities        []Activity
	Type              string // Company type (MEI, ME, LTDA, etc.)
	TenantId          uint   `validate:"required"`
	Schema            string
	Users             []User
}

func NewCompany(id int, legalName string, tradeName string, document string,
	stateRegistration string, address Address, partners []Partner,
	activities []Activity, ctype string, tenantId uint, schema string,
	users []User) *Company {
	return &Company{
		Id:                id,
		LegalName:         legalName,
		TradeName:         tradeName,
		Document:          document,
		StateRegistration: stateRegistration,
		Address:           address,
		Partners:          partners,
		Activities:        activities,
		Type:              ctype,
		TenantId:          tenantId,
		Schema:            schema,
		Users:             users,
	}
}
