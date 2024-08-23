package domain

type User struct {
	Id        int
	Name      string `validate:"required"`
	Email     string
	CPF       string
	Password  string
	Phone     string
	TenantId  int    `validate:"required"`
	Role      string `validate:"required,oneof=system_admin admin user client"`
	ProfileId int    // Opcional (orbigatorio para SystemUser) - usamos um ponteiro para permitir nulo
}

func NewUser(id int, name string, email string, cpf string, password string, phone string,
	tenantId int, role string, profileId int) *User {
	return &User{
		Id:        id,
		Name:      name,
		Email:     email,
		CPF:       cpf,
		Password:  password,
		Phone:     phone,
		TenantId:  tenantId,
		Role:      role,
		ProfileId: profileId,
	}
}
