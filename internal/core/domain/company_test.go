package domain

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestNewCompany(t *testing.T) {
	t.Run("endereço inválido", func(t *testing.T) {
		testCases := []struct {
			street  string
			number  string
			city    string
			state   string
			zipcode string
		}{
			{"", "123", "Cidade Grande", "SP", "12345-678"},
			{"Rua Exemplo", "", "Cidade Grande", "SP", "12345-678"},
			{"Rua Exemplo", "123", "", "SP", "12345-678"},
			{"Rua Exemplo", "123", "Cidade Grande", "", "12345-678"},
			{"Rua Exemplo", "123", "Cidade Grande", "SP", ""},
		}

		validate := validator.New()
		for _, tc := range testCases {
			address := Address{
				Street:     tc.street,
				Number:     tc.number,
				Complement: "Apto 456",
				District:   "Bairro Legal",
				City:       tc.city,
				State:      tc.state,
				ZipCode:    tc.zipcode,
			}

			company := NewCompany(1, "Nome Legal LTDA", "Nome Fantasia", "12345678901234",
				"123456789", address, []Partner{}, []Activity{}, "LTDA", 1, "public", []User{})
			err := validate.Struct(company)

			assert.Error(t, err)
			assert.True(t, IsRequiredFieldMissing(err))
		}
	})
}
