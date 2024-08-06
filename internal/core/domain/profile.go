package domain

type Profile struct {
	Id   int
	Name string `validate:"required"`
}

func NewProfile(id int, name string) *Profile {
	return &Profile{
		Id:   id,
		Name: name,
	}
}
