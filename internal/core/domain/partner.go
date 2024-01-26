package domain

type Partner struct {
	Participation float64 `validate:"required"`
	User          User
}

func NewPartner(participation float64, user User) *Partner {
	return &Partner{
		Participation: participation,
		User:          user,
	}
}
