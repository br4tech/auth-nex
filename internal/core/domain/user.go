package domain

type User struct {
	Name     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
	TenantID int    `validate:"required"`
	Roles    []Role
}

type Partner struct {
	Participation float64 `validate:"required"`
	User          User
}

func NewUser(name string, email string, password string,
	roles []Role, tenantID int) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Roles:    roles,
		TenantID: tenantID,
	}
}

func NewPartner(participation float64) *Partner {
	return &Partner{
		Participation: participation,
	}
}
