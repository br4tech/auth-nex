package domain

type User struct {
	Name      string `validate:"required"`
	Email     string `validate:"required"`
	Password  string `validate:"required"`
	TenantId  int    `validate:"required"`
	Companies []Company
	Roles     []Role
}

func NewUser(name string, email string, password string,
	roles []Role, tenantId int) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Roles:    roles,
		TenantId: tenantId,
	}
}
