package domain

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestNewAddress(t *testing.T) {
	t.Run("validação bem-sucedida", func(t *testing.T) {
		address := NewAddress("Rua Exemplo", "123", "Apto 456", "Bairro Legal", "Cidade Grande", "SP", "12345-678")
		validate := validator.New()
		err := validate.Struct(address)
		assert.NoError(t, err)
	})

	t.Run("campos obrigatórios ausentes", func(t *testing.T) {
		testCases := []struct {
			street     string
			number     string
			complement string
			district   string
			city       string
			state      string
			zipcode    string
		}{
			{"", "123", "Apto 456", "Bairro Legal", "Cidade Grande", "SP", "12345-678"},
			{"Rua Exemplo", "", "Apto 456", "Bairro Legal", "Cidade Grande", "SP", "12345-678"},
			{"Rua Exemplo", "123", "Apto 456", "Bairro Legal", "", "SP", "12345-678"},
			{"Rua Exemplo", "123", "Apto 456", "Bairro Legal", "Cidade Grande", "", "12345-678"},
			{"Rua Exemplo", "123", "Apto 456", "Bairro Legal", "Cidade Grande", "SP", ""},
		}

		validate := validator.New()
		for _, tc := range testCases {
			address := NewAddress(tc.street, tc.number, tc.complement, tc.district, tc.city, tc.state, tc.zipcode)
			err := validate.Struct(address)
			assert.Error(t, err)
			assert.True(t, IsRequiredFieldMissing(err))
		}
	})
}
