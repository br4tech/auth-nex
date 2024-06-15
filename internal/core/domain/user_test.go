package domain

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Run("validação bem-sucedida", func(t *testing.T) {
		user := NewUser("João Silva", "joao@example.com", "123.456.789-00", "senha123", 1, 2)
		validate := validator.New()
		err := validate.Struct(user)
		assert.NoError(t, err)
	})

	t.Run("campos obrigatórios ausentes", func(t *testing.T) {
		testCases := []struct {
			name      string
			email     string
			cpf       string
			password  string
			profileId int
			tenantId  int
		}{
			{"", "joao@example.com", "123.456.789-00", "senha123", 1, 2},
			{"João Silva", "", "123.456.789-00", "senha123", 1, 2},
			{"João Silva", "joao@example.com", "", "senha123", 1, 2},
			{"João Silva", "joao@example.com", "123.456.789-00", "", 1, 2},
			{"João Silva", "joao@example.com", "123.456.789-00", "senha123", 0, 2},
			{"João Silva", "joao@example.com", "123.456.789-00", "senha123", 1, 0},
		}

		validate := validator.New()
		for _, tc := range testCases {
			user := NewUser(tc.name, tc.email, tc.cpf, tc.password, tc.profileId, tc.tenantId)
			err := validate.Struct(user)
			assert.Error(t, err)
			assert.True(t, IsRequiredFieldMissing(err))
		}
	})

	t.Run("campos corretos", func(t *testing.T) {
		expectedUser := &User{
			Name:      "João Silva",
			Email:     "joao@example.com",
			CPF:       "123.456.789-00",
			Password:  "senha123",
			ProfileId: 1,
			TenantId:  2,
		}
		user := NewUser("João Silva", "joao@example.com", "123.456.789-00", "senha123", 1, 2)
		assert.Equal(t, expectedUser, user)
	})
}
