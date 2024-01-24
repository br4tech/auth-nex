package domain

type Address struct {
	Street     string `validate:"required"`
	Number     string `validate:"required"`
	Complement string
	District   string
	City       string `validate:"required"`
	State      string `validate:"required"`
	ZipCode    string `validate:"required"`
}

func NewAddress(street string, number string, complement string,
	district string, city string, state string, zipcode string) *Address {
	return &Address{
		Street:     street,
		Number:     number,
		Complement: complement,
		District:   district,
		City:       city,
		State:      state,
		ZipCode:    zipcode,
	}
}
