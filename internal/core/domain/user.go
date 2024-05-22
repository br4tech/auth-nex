package domain

type User struct {
	Name      string `validate:"required"`
	Email     string `validate:"required"`
	CPF       string `validate:"required"`
	Password  string `validate:"required"`
	TenantId  int    `validate:"required"`
	Companies []Company
	ProfileId int `validate:"required"`
}

func NewUser(name string, email string, cpf string, password string,
	profileId int, tenantId int) *User {
	return &User{
		Name:      name,
		Email:     email,
		CPF:       cpf,
		Password:  password,
		ProfileId: profileId,
		TenantId:  tenantId,
	}
}
