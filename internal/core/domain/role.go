package domain

type Role struct {
	Name string `validate:"required"`
}

func NewRole(name string) *Role {
	return &Role{
		Name: name,
	}
}
