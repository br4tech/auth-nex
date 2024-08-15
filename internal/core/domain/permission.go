package domain

type Permission struct {
	Id   int    `validate:"required"`
	Name string `validate:"required"`
}

func NewPermission(id int, name string) *Permission {
	return &Permission{
		Id:   id,
		Name: name,
	}
}
