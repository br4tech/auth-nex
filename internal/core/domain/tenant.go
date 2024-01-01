package domain

type Tenant struct {
	Name string `validate:"required"`
}

func NewTenant(name string) *Tenant {
	return &Tenant{
		Name: name,
	}
}
