package domain

type User struct {
	Name     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
	Roles    []Role
	TenantID int `validate:"required"`
}

func NewUser(
	name string,
	email string,
	password string,
	roles []Role,
	tenantID int) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Roles:    roles,
		TenantID: tenantID,
	}
}
