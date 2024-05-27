package domain

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestNewTenant(t *testing.T) {
	t.Run("validação bem-sucedida", func(t *testing.T) {
		tenant := NewTenant("Tenant Exemplo", []Company{}, []User{})
		validate := validator.New()
		err := validate.Struct(tenant)
		assert.NoError(t, err)
	})

	t.Run("nome ausente", func(t *testing.T) {
		tenant := NewTenant("", []Company{}, []User{})
		validate := validator.New()
		err := validate.Struct(tenant)
		assert.Error(t, err)
		assert.True(t, IsRequiredFieldMissing(err))
	})

	t.Run("campos corretos", func(t *testing.T) {
		companies := []Company{{LegalName: "Empresa A"}, {LegalName: "Empresa B"}}
		users := []User{{Name: "João"}, {Name: "Maria"}}
		expectedTenant := &Tenant{
			Name:      "Tenant Exemplo",
			Companies: companies,
			Users:     users,
		}
		tenant := NewTenant("Tenant Exemplo", companies, users)
		assert.Equal(t, expectedTenant, tenant)
	})
}
