package domain

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestNewPartner(t *testing.T) {
	// t.Run("validação bem-sucedida", func(t *testing.T) {

	// 	partner := NewPartner(33.33, User{}, Company{}) // User preenchido

	// 	validate := validator.New()
	// 	err := validate.Struct(partner)
	// 	assert.NoError(t, err)
	// })

	t.Run("campos obrigatórios ausentes", func(t *testing.T) {
		testCases := []struct {
			participation float64
			user          User
			company       Company
		}{
			{0, User{}, Company{}},                 // Participation ausente
			{33.33, User{}, Company{}},             // User ausente
			{33.33, User{Name: "João"}, Company{}}, // Company ausente
		}

		validate := validator.New()
		for _, tc := range testCases {
			partner := NewPartner(tc.participation, tc.user)
			partner.Company = tc.company
			err := validate.Struct(partner)
			assert.Error(t, err)

			if tc.participation == 0 {
				assert.True(t, IsRequiredFieldMissing(err))
			} else if !cmp.Equal(tc.company, Company{}) {
				assert.True(t, err.Error() == "Key: 'Partner.Company' Error:Field validation for 'Company' failed on the 'required' tag")
			}
		}
	})

	t.Run("campos corretos", func(t *testing.T) {
		user := User{Name: "João"}
		company := Company{LegalName: "Empresa"}
		expectedPartner := &Partner{
			Participation: 33.33,
			User:          user,
			Company:       company,
		}

		partner := NewPartner(33.33, user)
		partner.Company = company

		assert.Equal(t, expectedPartner, partner)
	})
}
