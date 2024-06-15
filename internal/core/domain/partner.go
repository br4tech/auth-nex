package domain

type Partner struct {
	Participation float64 `validate:"required"`
	UserId        int     `validate:"required"`
	User          User
	CompanyId     int
	Company       Company
}

func NewPartner(participation float64, user User) *Partner {
	return &Partner{
		Participation: participation,
		User:          user,
	}
}
