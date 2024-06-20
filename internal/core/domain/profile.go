package domain

type Profile struct {
	Name string `validate:"required"`
}

func NewProfile(name string) *Profile {
	return &Profile{
		Name: name,
	}
}
