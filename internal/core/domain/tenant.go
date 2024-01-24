package domain

type Tenant struct {
	Name      string `validate:"required"`
	Companies []Company
	Users     []User
}

func NewTenant(name string, companies []Company, users []User) *Tenant {
	return &Tenant{
		Name:      name,
		Companies: companies,
		Users:     users,
	}
}
