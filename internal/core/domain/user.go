package domain

type User struct {
	Name  string `validate:"required"`
	Email string `validate:"required"`
	Roles []Role
}

func NewUser(name string, email string, roles []Role) *User {
	return &User{
		Name:  name,
		Email: email,
		Roles: roles,
	}
}
