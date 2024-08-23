package domain

type Activity struct {
	CNAE        string `validate:"required"`
	Description string `validate:"required"`
}

func NewActivity(cnae string, description string) *Activity {
	return &Activity{
		CNAE:        cnae,
		Description: description,
	}
}
