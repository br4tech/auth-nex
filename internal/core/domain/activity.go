package domain

type Activity struct {
	CNAE        string `validate:"required"`
	Description string `validate:"required"`
	CompanyID   uint   `validate:"required"`
}

func NewActivity(cnae string, description string, company_id int) *Activity {
	return &Activity{
		CNAE:        cnae,
		Description: description,
		CompanyID:   uint(company_id),
	}
}
