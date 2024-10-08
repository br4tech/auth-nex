package domain

type Company struct {
	Id                int
	LegalName         string `validate:"required"`
	TradeName         string `validate:"required"`
	Document          string `validate:"required"`
	StateRegistration string `validate:"required"`
	Type              string // Company type (MEI, ME, LTDA, etc.)
	TenantId          int    `validate:"required"`
	Schema            string
	ParentCompanyId   int
	Address           Address
	Users             []User
	Partners          []Partner
	Activities        []Activity
}

func NewCompany(legalName string, tradeName string, document string,
	stateRegistration string, address Address, partners []Partner, parentCompanyId int,
	activities []Activity, ctype string, tenantId int, schema string,
	users []User) *Company {
	return &Company{
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
		ParentCompanyId:   parentCompanyId,
	}
}
