package domain

type User struct {
	Id        int
	Name      string `validate:"required"`
	Email     string
	CPF       *string // Opcional (orbigatorio para SystemUser) - usamos um ponteiro para permitir nulo
	Password  string  // Opcional (orbigatorio para SystemUser) - usamos um ponteiro para permitir nulo
	Phone     string  // Opcional (mas obrigatório para ClientUser, a menos que Email esteja presente)
	TenantId  int     `validate:"required"`
	Role      string  `validate:"required,oneof=system client"`
	ProfileId *int    // Opcional (orbigatorio para SystemUser) - usamos um ponteiro para permitir nulo
}

func NewUser(id int, name string, email string, cpf *string, password string, phone string,
	tenantId int, role string, profileId *int) *User {
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
