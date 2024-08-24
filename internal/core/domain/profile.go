package domain

type Profile struct {
	Id         int
	Name       string `validate:"required"`
	TenantId   int
	Permisions []Permission
}

func NewProfile(id int, name string, tenantId int, permissions []Permission) *Profile {
	return &Profile{
		Id:         id,
		Name:       name,
		TenantId:   tenantId,
		Permisions: permissions,
	}
}
